package agent

import types "github.com/sensu/sensu-go/api/core/v2"

// A Transformer handles transforming Sensu metrics to other output metric formats
type Transformer interface {
	// Transform transforms a metric in a different output metric format to Sensu Metric
	// Format
	Transform() []*types.MetricPoint
}
