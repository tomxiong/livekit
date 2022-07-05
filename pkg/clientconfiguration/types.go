package clientconfiguration

import (
	"github.com/tomxiong/protocol/livekit"
)

type ClientConfigurationManager interface {
	GetConfiguration(clientInfo *livekit.ClientInfo) *livekit.ClientConfiguration
}
