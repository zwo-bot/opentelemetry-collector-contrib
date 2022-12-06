// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestDefaultMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	mb := NewMetricsBuilder(DefaultMetricsSettings(), componenttest.NewNopReceiverCreateSettings(), WithStartTime(start))
	enabledMetrics := make(map[string]bool)

	enabledMetrics["zookeeper.connection.active"] = true
	mb.RecordZookeeperConnectionActiveDataPoint(ts, 1)

	enabledMetrics["zookeeper.data_tree.ephemeral_node.count"] = true
	mb.RecordZookeeperDataTreeEphemeralNodeCountDataPoint(ts, 1)

	enabledMetrics["zookeeper.data_tree.size"] = true
	mb.RecordZookeeperDataTreeSizeDataPoint(ts, 1)

	enabledMetrics["zookeeper.file_descriptor.limit"] = true
	mb.RecordZookeeperFileDescriptorLimitDataPoint(ts, 1)

	enabledMetrics["zookeeper.file_descriptor.open"] = true
	mb.RecordZookeeperFileDescriptorOpenDataPoint(ts, 1)

	enabledMetrics["zookeeper.follower.count"] = true
	mb.RecordZookeeperFollowerCountDataPoint(ts, 1, AttributeState(1))

	enabledMetrics["zookeeper.fsync.exceeded_threshold.count"] = true
	mb.RecordZookeeperFsyncExceededThresholdCountDataPoint(ts, 1)

	enabledMetrics["zookeeper.latency.avg"] = true
	mb.RecordZookeeperLatencyAvgDataPoint(ts, 1)

	enabledMetrics["zookeeper.latency.max"] = true
	mb.RecordZookeeperLatencyMaxDataPoint(ts, 1)

	enabledMetrics["zookeeper.latency.min"] = true
	mb.RecordZookeeperLatencyMinDataPoint(ts, 1)

	enabledMetrics["zookeeper.packet.count"] = true
	mb.RecordZookeeperPacketCountDataPoint(ts, 1, AttributeDirection(1))

	enabledMetrics["zookeeper.request.active"] = true
	mb.RecordZookeeperRequestActiveDataPoint(ts, 1)

	enabledMetrics["zookeeper.sync.pending"] = true
	mb.RecordZookeeperSyncPendingDataPoint(ts, 1)

	enabledMetrics["zookeeper.watch.count"] = true
	mb.RecordZookeeperWatchCountDataPoint(ts, 1)

	enabledMetrics["zookeeper.znode.count"] = true
	mb.RecordZookeeperZnodeCountDataPoint(ts, 1)

	metrics := mb.Emit()

	assert.Equal(t, 1, metrics.ResourceMetrics().Len())
	sm := metrics.ResourceMetrics().At(0).ScopeMetrics()
	assert.Equal(t, 1, sm.Len())
	ms := sm.At(0).Metrics()
	assert.Equal(t, len(enabledMetrics), ms.Len())
	seenMetrics := make(map[string]bool)
	for i := 0; i < ms.Len(); i++ {
		assert.True(t, enabledMetrics[ms.At(i).Name()])
		seenMetrics[ms.At(i).Name()] = true
	}
	assert.Equal(t, len(enabledMetrics), len(seenMetrics))
}

func TestAllMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	metricsSettings := MetricsSettings{
		ZookeeperConnectionActive:            MetricSettings{Enabled: true},
		ZookeeperDataTreeEphemeralNodeCount:  MetricSettings{Enabled: true},
		ZookeeperDataTreeSize:                MetricSettings{Enabled: true},
		ZookeeperFileDescriptorLimit:         MetricSettings{Enabled: true},
		ZookeeperFileDescriptorOpen:          MetricSettings{Enabled: true},
		ZookeeperFollowerCount:               MetricSettings{Enabled: true},
		ZookeeperFsyncExceededThresholdCount: MetricSettings{Enabled: true},
		ZookeeperLatencyAvg:                  MetricSettings{Enabled: true},
		ZookeeperLatencyMax:                  MetricSettings{Enabled: true},
		ZookeeperLatencyMin:                  MetricSettings{Enabled: true},
		ZookeeperPacketCount:                 MetricSettings{Enabled: true},
		ZookeeperRequestActive:               MetricSettings{Enabled: true},
		ZookeeperSyncPending:                 MetricSettings{Enabled: true},
		ZookeeperWatchCount:                  MetricSettings{Enabled: true},
		ZookeeperZnodeCount:                  MetricSettings{Enabled: true},
	}
	observedZapCore, observedLogs := observer.New(zap.WarnLevel)
	settings := componenttest.NewNopReceiverCreateSettings()
	settings.Logger = zap.New(observedZapCore)
	mb := NewMetricsBuilder(metricsSettings, settings, WithStartTime(start))

	assert.Equal(t, 0, observedLogs.Len())

	mb.RecordZookeeperConnectionActiveDataPoint(ts, 1)
	mb.RecordZookeeperDataTreeEphemeralNodeCountDataPoint(ts, 1)
	mb.RecordZookeeperDataTreeSizeDataPoint(ts, 1)
	mb.RecordZookeeperFileDescriptorLimitDataPoint(ts, 1)
	mb.RecordZookeeperFileDescriptorOpenDataPoint(ts, 1)
	mb.RecordZookeeperFollowerCountDataPoint(ts, 1, AttributeState(1))
	mb.RecordZookeeperFsyncExceededThresholdCountDataPoint(ts, 1)
	mb.RecordZookeeperLatencyAvgDataPoint(ts, 1)
	mb.RecordZookeeperLatencyMaxDataPoint(ts, 1)
	mb.RecordZookeeperLatencyMinDataPoint(ts, 1)
	mb.RecordZookeeperPacketCountDataPoint(ts, 1, AttributeDirection(1))
	mb.RecordZookeeperRequestActiveDataPoint(ts, 1)
	mb.RecordZookeeperSyncPendingDataPoint(ts, 1)
	mb.RecordZookeeperWatchCountDataPoint(ts, 1)
	mb.RecordZookeeperZnodeCountDataPoint(ts, 1)

	metrics := mb.Emit(WithServerState("attr-val"), WithZkVersion("attr-val"))

	assert.Equal(t, 1, metrics.ResourceMetrics().Len())
	rm := metrics.ResourceMetrics().At(0)
	attrCount := 0
	attrCount++
	attrVal, ok := rm.Resource().Attributes().Get("server.state")
	assert.True(t, ok)
	assert.EqualValues(t, "attr-val", attrVal.Str())
	attrCount++
	attrVal, ok = rm.Resource().Attributes().Get("zk.version")
	assert.True(t, ok)
	assert.EqualValues(t, "attr-val", attrVal.Str())
	assert.Equal(t, attrCount, rm.Resource().Attributes().Len())

	assert.Equal(t, 1, rm.ScopeMetrics().Len())
	ms := rm.ScopeMetrics().At(0).Metrics()
	allMetricsCount := reflect.TypeOf(MetricsSettings{}).NumField()
	assert.Equal(t, allMetricsCount, ms.Len())
	validatedMetrics := make(map[string]struct{})
	for i := 0; i < ms.Len(); i++ {
		switch ms.At(i).Name() {
		case "zookeeper.connection.active":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of active clients connected to a ZooKeeper server.", ms.At(i).Description())
			assert.Equal(t, "{connections}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["zookeeper.connection.active"] = struct{}{}
		case "zookeeper.data_tree.ephemeral_node.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of ephemeral nodes that a ZooKeeper server has in its data tree.", ms.At(i).Description())
			assert.Equal(t, "{nodes}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["zookeeper.data_tree.ephemeral_node.count"] = struct{}{}
		case "zookeeper.data_tree.size":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Size of data in bytes that a ZooKeeper server has in its data tree.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["zookeeper.data_tree.size"] = struct{}{}
		case "zookeeper.file_descriptor.limit":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Maximum number of file descriptors that a ZooKeeper server can open.", ms.At(i).Description())
			assert.Equal(t, "{file_descriptors}", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["zookeeper.file_descriptor.limit"] = struct{}{}
		case "zookeeper.file_descriptor.open":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of file descriptors that a ZooKeeper server has open.", ms.At(i).Description())
			assert.Equal(t, "{file_descriptors}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["zookeeper.file_descriptor.open"] = struct{}{}
		case "zookeeper.follower.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of followers. Only exposed by the leader.", ms.At(i).Description())
			assert.Equal(t, "{followers}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("state")
			assert.True(t, ok)
			assert.Equal(t, "synced", attrVal.Str())
			validatedMetrics["zookeeper.follower.count"] = struct{}{}
		case "zookeeper.fsync.exceeded_threshold.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of times fsync duration has exceeded warning threshold.", ms.At(i).Description())
			assert.Equal(t, "{events}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["zookeeper.fsync.exceeded_threshold.count"] = struct{}{}
		case "zookeeper.latency.avg":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Average time in milliseconds for requests to be processed.", ms.At(i).Description())
			assert.Equal(t, "ms", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["zookeeper.latency.avg"] = struct{}{}
		case "zookeeper.latency.max":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Maximum time in milliseconds for requests to be processed.", ms.At(i).Description())
			assert.Equal(t, "ms", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["zookeeper.latency.max"] = struct{}{}
		case "zookeeper.latency.min":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Minimum time in milliseconds for requests to be processed.", ms.At(i).Description())
			assert.Equal(t, "ms", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["zookeeper.latency.min"] = struct{}{}
		case "zookeeper.packet.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of ZooKeeper packets received or sent by a server.", ms.At(i).Description())
			assert.Equal(t, "{packets}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("direction")
			assert.True(t, ok)
			assert.Equal(t, "received", attrVal.Str())
			validatedMetrics["zookeeper.packet.count"] = struct{}{}
		case "zookeeper.request.active":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of currently executing requests.", ms.At(i).Description())
			assert.Equal(t, "{requests}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["zookeeper.request.active"] = struct{}{}
		case "zookeeper.sync.pending":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of pending syncs from the followers. Only exposed by the leader.", ms.At(i).Description())
			assert.Equal(t, "{syncs}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["zookeeper.sync.pending"] = struct{}{}
		case "zookeeper.watch.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of watches placed on Z-Nodes on a ZooKeeper server.", ms.At(i).Description())
			assert.Equal(t, "{watches}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["zookeeper.watch.count"] = struct{}{}
		case "zookeeper.znode.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of z-nodes that a ZooKeeper server has in its data tree.", ms.At(i).Description())
			assert.Equal(t, "{znodes}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["zookeeper.znode.count"] = struct{}{}
		}
	}
	assert.Equal(t, allMetricsCount, len(validatedMetrics))
}

func TestNoMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	metricsSettings := MetricsSettings{
		ZookeeperConnectionActive:            MetricSettings{Enabled: false},
		ZookeeperDataTreeEphemeralNodeCount:  MetricSettings{Enabled: false},
		ZookeeperDataTreeSize:                MetricSettings{Enabled: false},
		ZookeeperFileDescriptorLimit:         MetricSettings{Enabled: false},
		ZookeeperFileDescriptorOpen:          MetricSettings{Enabled: false},
		ZookeeperFollowerCount:               MetricSettings{Enabled: false},
		ZookeeperFsyncExceededThresholdCount: MetricSettings{Enabled: false},
		ZookeeperLatencyAvg:                  MetricSettings{Enabled: false},
		ZookeeperLatencyMax:                  MetricSettings{Enabled: false},
		ZookeeperLatencyMin:                  MetricSettings{Enabled: false},
		ZookeeperPacketCount:                 MetricSettings{Enabled: false},
		ZookeeperRequestActive:               MetricSettings{Enabled: false},
		ZookeeperSyncPending:                 MetricSettings{Enabled: false},
		ZookeeperWatchCount:                  MetricSettings{Enabled: false},
		ZookeeperZnodeCount:                  MetricSettings{Enabled: false},
	}
	observedZapCore, observedLogs := observer.New(zap.WarnLevel)
	settings := componenttest.NewNopReceiverCreateSettings()
	settings.Logger = zap.New(observedZapCore)
	mb := NewMetricsBuilder(metricsSettings, settings, WithStartTime(start))

	assert.Equal(t, 0, observedLogs.Len())
	mb.RecordZookeeperConnectionActiveDataPoint(ts, 1)
	mb.RecordZookeeperDataTreeEphemeralNodeCountDataPoint(ts, 1)
	mb.RecordZookeeperDataTreeSizeDataPoint(ts, 1)
	mb.RecordZookeeperFileDescriptorLimitDataPoint(ts, 1)
	mb.RecordZookeeperFileDescriptorOpenDataPoint(ts, 1)
	mb.RecordZookeeperFollowerCountDataPoint(ts, 1, AttributeState(1))
	mb.RecordZookeeperFsyncExceededThresholdCountDataPoint(ts, 1)
	mb.RecordZookeeperLatencyAvgDataPoint(ts, 1)
	mb.RecordZookeeperLatencyMaxDataPoint(ts, 1)
	mb.RecordZookeeperLatencyMinDataPoint(ts, 1)
	mb.RecordZookeeperPacketCountDataPoint(ts, 1, AttributeDirection(1))
	mb.RecordZookeeperRequestActiveDataPoint(ts, 1)
	mb.RecordZookeeperSyncPendingDataPoint(ts, 1)
	mb.RecordZookeeperWatchCountDataPoint(ts, 1)
	mb.RecordZookeeperZnodeCountDataPoint(ts, 1)

	metrics := mb.Emit()

	assert.Equal(t, 0, metrics.ResourceMetrics().Len())
}
