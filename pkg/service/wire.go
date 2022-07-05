//go:build wireinject
// +build wireinject

package service

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/tomxiong/protocol/auth"
	"github.com/tomxiong/protocol/egress"
	"github.com/tomxiong/protocol/livekit"
	"github.com/tomxiong/protocol/logger"
	"github.com/tomxiong/protocol/utils"
	"github.com/tomxiong/protocol/webhook"

	"github.com/livekit/livekit-server/pkg/clientconfiguration"
	"github.com/livekit/livekit-server/pkg/config"
	"github.com/livekit/livekit-server/pkg/routing"
	"github.com/livekit/livekit-server/pkg/telemetry"
	redisClient "github.com/tomxiong/protocol/utils/redis"
)

func InitializeServer(conf *config.Config, currentNode routing.LocalNode) (*LivekitServer, error) {
	wire.Build(
		getNodeID,
		createRedisClient,
		createMessageBus,
		createStore,
		wire.Bind(new(ServiceStore), new(ObjectStore)),
		createKeyProvider,
		createWebhookNotifier,
		createClientConfiguration,
		routing.CreateRouter,
		getRoomConf,
		wire.Bind(new(routing.MessageRouter), new(routing.Router)),
		wire.Bind(new(livekit.RoomService), new(*RoomService)),
		telemetry.NewAnalyticsService,
		telemetry.NewTelemetryService,
		egress.NewRedisRPCClient,
		NewEgressService,
		NewRecordingService,
		NewRoomAllocator,
		NewRoomService,
		NewRTCService,
		NewLocalRoomManager,
		newTurnAuthHandler,
		NewTurnServer,
		NewLivekitServer,
	)
	return &LivekitServer{}, nil
}

func InitializeRouter(conf *config.Config, currentNode routing.LocalNode) (routing.Router, error) {
	wire.Build(
		createRedisClient,
		routing.CreateRouter,
	)

	return nil, nil
}

func getNodeID(currentNode routing.LocalNode) livekit.NodeID {
	return livekit.NodeID(currentNode.Id)
}

func createKeyProvider(conf *config.Config) (auth.KeyProvider, error) {
	// prefer keyfile if set
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

func createRedisClient(conf *config.Config) (redisClient.RedisClient, error) {
	if !conf.HasRedis() {
		return nil, nil
	}

	var rc redisClient.RedisClient
	var tlsConfig *tls.Config

	if conf.Redis.UseTLS {
		tlsConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	if conf.UseCluster() {
		logger.Infow("using multi-node routing via redis", "cluster", true, "addr", conf.Redis.ClusterAddresses)
		rcOptions := &redis.ClusterOptions{
			Addrs:     conf.Redis.ClusterAddresses,
			Password:  conf.Redis.Password,
			TLSConfig: tlsConfig,
		}
		rc = redis.NewClusterClient(rcOptions)
	} else if conf.UseSentinel() {
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
		logger.Infow("using single-node redis", "single", true, "addr", conf.Redis.Address)
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

func createMessageBus(rc redisClient.RedisClient) utils.MessageBus {
	if rc == nil {
		return nil
	}
	return utils.NewRedisMessageBus(rc)
}

func createStore(rc redisClient.RedisClient) ObjectStore {
	if rc != nil {
		return NewRedisStore(rc)
	}
	return NewLocalStore()
}

func createClientConfiguration() clientconfiguration.ClientConfigurationManager {
	return clientconfiguration.NewStaticClientConfigurationManager(clientconfiguration.StaticConfigurations)
}

func getRoomConf(config *config.Config) config.RoomConfig {
	return config.Room
}
