package monitoring

// MetricType holds tsdb metric type.
type MetricType string

const (
	// MetricTypeCounter holds tsdb counter metric type.
	MetricTypeCounter MetricType = "counter"
	// MetricTypeGauge holds tsdb gauge metric type.
	MetricTypeGauge MetricType = "gauge"
)

// Event represents the metric event.
type Event interface {
	Name() string
	Value() interface{}
	Labels() map[string]string
	MetricType() MetricType
}

// Events is custom type for []Event
type Events []Event
