package prometheus

import (
	"time"

	"github.com/mackerelio/go-osstat/loadavg"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/livekit/livekit-server/pkg/config"
	"github.com/tomxiong/protocol/livekit"
	"github.com/tomxiong/protocol/utils"
)

const (
	livekitNamespace string = "livekit"
)

var (
	MessageCounter          *prometheus.CounterVec
	ServiceOperationCounter *prometheus.CounterVec
)

func init() {
	nodeID, _ := utils.LocalNodeID()
	MessageCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   livekitNamespace,
			Subsystem:   "node",
			Name:        "messages",
			ConstLabels: prometheus.Labels{"node_id": nodeID},
		},
		[]string{"type", "status"},
	)

	ServiceOperationCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   livekitNamespace,
			Subsystem:   "node",
			Name:        "service_operation",
			ConstLabels: prometheus.Labels{"node_id": nodeID},
		},
		[]string{"type", "status", "error_type"},
	)

	prometheus.MustRegister(MessageCounter)
	prometheus.MustRegister(ServiceOperationCounter)

	initPacketStats(nodeID)
	initRoomStats(nodeID)
}

func GetUpdatedNodeStats(prev *livekit.NodeStats, prevAverage *livekit.NodeStats) (*livekit.NodeStats, bool, error) {
	loadAvg, err := loadavg.Get()
	if err != nil {
		return nil, false, err
	}

	cpuLoad, numCPUs, err := getCPUStats()
	if err != nil {
		return nil, false, err
	}

	bytesInNow := bytesIn.Load()
	bytesOutNow := bytesOut.Load()
	packetsInNow := packetsIn.Load()
	packetsOutNow := packetsOut.Load()
	nackTotalNow := nackTotal.Load()
	retransmitBytesNow := retransmitBytes.Load()
	retransmitPacketsNow := retransmitPackets.Load()
	participantJoinNow := participantJoin.Load()

	updatedAt := time.Now().Unix()
	elapsed := updatedAt - prevAverage.UpdatedAt
	// include sufficient buffer to be sure a stats update had taken place
	computeAverage := elapsed > int64(config.StatsUpdateInterval.Seconds()+2)
	if bytesInNow != prevAverage.BytesIn ||
		bytesOutNow != prevAverage.BytesOut ||
		packetsInNow != prevAverage.PacketsIn ||
		packetsOutNow != prevAverage.PacketsOut ||
		retransmitBytesNow != prevAverage.RetransmitBytesOut ||
		retransmitPacketsNow != prevAverage.RetransmitPacketsOut {
		computeAverage = true
	}

	stats := &livekit.NodeStats{
		StartedAt:                  prev.StartedAt,
		UpdatedAt:                  updatedAt,
		NumRooms:                   roomTotal.Load(),
		NumClients:                 participantTotal.Load(),
		NumTracksIn:                trackPublishedTotal.Load(),
		NumTracksOut:               trackSubscribedTotal.Load(),
		BytesIn:                    bytesInNow,
		BytesOut:                   bytesOutNow,
		PacketsIn:                  packetsInNow,
		PacketsOut:                 packetsOutNow,
		RetransmitBytesOut:         retransmitBytesNow,
		RetransmitPacketsOut:       retransmitPacketsNow,
		NackTotal:                  nackTotalNow,
		ParticipantJoin:            participantJoinNow,
		BytesInPerSec:              prevAverage.BytesInPerSec,
		BytesOutPerSec:             prevAverage.BytesOutPerSec,
		PacketsInPerSec:            prevAverage.PacketsInPerSec,
		PacketsOutPerSec:           prevAverage.PacketsOutPerSec,
		RetransmitBytesOutPerSec:   prevAverage.RetransmitBytesOutPerSec,
		RetransmitPacketsOutPerSec: prevAverage.RetransmitPacketsOutPerSec,
		NackPerSec:                 prevAverage.NackPerSec,
		ParticipantJoinPerSec:      prevAverage.ParticipantJoinPerSec,
		NumCpus:                    numCPUs,
		CpuLoad:                    cpuLoad,
		LoadAvgLast1Min:            float32(loadAvg.Loadavg1),
		LoadAvgLast5Min:            float32(loadAvg.Loadavg5),
		LoadAvgLast15Min:           float32(loadAvg.Loadavg15),
	}

	// update stats
	if computeAverage {
		stats.BytesInPerSec = perSec(prevAverage.BytesIn, bytesInNow, elapsed)
		stats.BytesOutPerSec = perSec(prevAverage.BytesOut, bytesOutNow, elapsed)
		stats.PacketsInPerSec = perSec(prevAverage.PacketsIn, packetsInNow, elapsed)
		stats.PacketsOutPerSec = perSec(prevAverage.PacketsOut, packetsOutNow, elapsed)
		stats.RetransmitBytesOutPerSec = perSec(prevAverage.RetransmitBytesOut, retransmitBytesNow, elapsed)
		stats.RetransmitPacketsOutPerSec = perSec(prevAverage.RetransmitPacketsOut, retransmitPacketsNow, elapsed)
		stats.NackPerSec = perSec(prevAverage.NackTotal, nackTotalNow, elapsed)
		stats.ParticipantJoinPerSec = perSec(prevAverage.ParticipantJoin, participantJoinNow, elapsed)
	}

	return stats, computeAverage, nil
}

func perSec(prev, curr uint64, secs int64) float32 {
	return float32(curr-prev) / float32(secs)
}
