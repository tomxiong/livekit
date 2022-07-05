// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/livekit/livekit-server/pkg/clientconfiguration"
	"github.com/livekit/livekit-server/pkg/config"
	"github.com/livekit/livekit-server/pkg/routing"
	"github.com/livekit/livekit-server/pkg/telemetry"
	"github.com/pkg/errors"
	"github.com/tomxiong/protocol/auth"
	"github.com/tomxiong/protocol/egress"
	"github.com/tomxiong/protocol/livekit"
	"github.com/tomxiong/protocol/logger"
	"github.com/tomxiong/protocol/utils"
	"github.com/tomxiong/protocol/webhook"
	"gopkg.in/yaml.v3"
	"os"
)

import (
	_ "net/http/pprof"
)

// Injectors from wire.go:

func InitializeServer(conf *config.Config, currentNode routing.LocalNode) (*LivekitServer, error) {
	client, err := createRedisClient(conf)
	if err != nil {
		return nil, err
	}
	router := routing.CreateRouter(client, currentNode)
	objectStore := createStore(client)
	roomAllocator, err := NewRoomAllocator(conf, router, objectStore)
	if err != nil {
		return nil, err
	}
	roomConfig := getRoomConf(conf)
	roomService, err := NewRoomService(roomAllocator, objectStore, router, roomConfig)
	if err != nil {
		return nil, err
	}
	nodeID := getNodeID(currentNode)
	rpcClient := egress.NewRedisRPCClient(nodeID, client)
	keyProvider, err := createKeyProvider(conf)
	if err != nil {
		return nil, err
	}
	notifier, err := createWebhookNotifier(conf, keyProvider)
	if err != nil {
		return nil, err
	}
	analyticsService := telemetry.NewAnalyticsService(conf, currentNode)
	telemetryService := telemetry.NewTelemetryService(notifier, analyticsService)
	egressService := NewEgressService(rpcClient, objectStore, roomService, telemetryService)
	messageBus := createMessageBus(client)
	recordingService := NewRecordingService(messageBus, telemetryService)
	rtcService := NewRTCService(conf, roomAllocator, objectStore, router, currentNode)
	clientConfigurationManager := createClientConfiguration()
	roomManager, err := NewLocalRoomManager(conf, objectStore, currentNode, router, telemetryService, clientConfigurationManager)
	if err != nil {
		return nil, err
	}
	authHandler := newTurnAuthHandler(objectStore)
	server, err := NewTurnServer(conf, authHandler)
	if err != nil {
		return nil, err
	}
	livekitServer, err := NewLivekitServer(conf, roomService, egressService, recordingService, rtcService, keyProvider, router, roomManager, server, currentNode)
	if err != nil {
		return nil, err
	}
	return livekitServer, nil
}

func InitializeRouter(conf *config.Config, currentNode routing.LocalNode) (routing.Router, error) {
	client, err := createRedisClient(conf)
	if err != nil {
		return nil, err
	}
	router := routing.CreateRouter(client, currentNode)
	return router, nil
}

// wire.go:

func getNodeID(currentNode routing.LocalNode) livekit.NodeID {
	return livekit.NodeID(currentNode.Id)
}

func createKeyProvider(conf *config.Config) (auth.KeyProvider, error) {

	if conf.KeyFile != "" {
		if st, err := os.Stat(conf.KeyFile); err != nil {
			return nil, err
		} else if st.Mode().Perm() != 0600 {
			return nil, fmt.Errorf("key file must have permission set to 600")
		}
		f, err := os.Open(conf.KeyFile)
		if err != nil {
			return nil, err
		}
		defer func() {
			_ = f.Close()
		}()
		decoder := yaml.NewDecoder(f)
		if err = decoder.Decode(conf.Keys); err != nil {
			return nil, err
		}
	}

	if len(conf.Keys) == 0 {
		return nil, errors.New("one of key-file or keys must be provided in order to support a secure installation")
	}

	return auth.NewFileBasedKeyProviderFromMap(conf.Keys), nil
}

func createWebhookNotifier(conf *config.Config, provider auth.KeyProvider) (webhook.Notifier, error) {
	wc := conf.WebHook
	if len(wc.URLs) == 0 {
		return nil, nil
	}
	secret := provider.GetSecret(wc.APIKey)
	if secret == "" {
		return nil, ErrWebHookMissingAPIKey
	}

	return webhook.NewNotifier(wc.APIKey, secret, wc.URLs), nil
}

func createRedisClient(conf *config.Config) (*redis.Client, error) {
	if !conf.HasRedis() {
		return nil, nil
	}

	var rc *redis.Client
	var tlsConfig *tls.Config

	if conf.Redis.UseTLS {
		tlsConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	if conf.UseSentinel() {
		logger.Infow("using multi-node routing via redis", "sentinel", true, "addr", conf.Redis.SentinelAddresses, "masterName", conf.Redis.MasterName)
		rcOptions := &redis.FailoverOptions{
			SentinelAddrs:    conf.Redis.SentinelAddresses,
			SentinelUsername: conf.Redis.SentinelUsername,
			SentinelPassword: conf.Redis.SentinelPassword,
			MasterName:       conf.Redis.MasterName,
			Username:         conf.Redis.Username,
			Password:         conf.Redis.Password,
			DB:               conf.Redis.DB,
			TLSConfig:        tlsConfig,
		}
		rc = redis.NewFailoverClient(rcOptions)
	} else {
		logger.Infow("using multi-node routing via redis", "sentinel", false, "addr", conf.Redis.Address)
		rcOptions := &redis.Options{
			Addr:      conf.Redis.Address,
			Username:  conf.Redis.Username,
			Password:  conf.Redis.Password,
			DB:        conf.Redis.DB,
			TLSConfig: tlsConfig,
		}
		rc = redis.NewClient(rcOptions)
	}

	if err := rc.Ping(context.Background()).Err(); err != nil {
		err = errors.Wrap(err, "unable to connect to redis")
		return nil, err
	}

	return rc, nil
}

func createMessageBus(rc *redis.Client) utils.MessageBus {
	if rc == nil {
		return nil
	}
	return utils.NewRedisMessageBus(rc)
}

func createStore(rc *redis.Client) ObjectStore {
	if rc != nil {
		return NewRedisStore(rc)
	}
	return NewLocalStore()
}

func createClientConfiguration() clientconfiguration.ClientConfigurationManager {
	return clientconfiguration.NewStaticClientConfigurationManager(clientconfiguration.StaticConfigurations)
}

func getRoomConf(config2 *config.Config) config.RoomConfig {
	return config2.Room
}
