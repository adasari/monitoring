package monitoring

// IntGauge is metric of gauge type.
type IntGauge struct {
	id        string
	projectID string
	name      string
	val       interface{}
	labels    map[string]string
}

func newIntGauge(projectID string, name string, labels map[string]string) *IntGauge {
	return &IntGauge{
		projectID: projectID,
		name:      name,
		labels:    labels,
	}
}

func (g *IntGauge) Name() string              { return g.name }
func (g *IntGauge) Value() interface{}        { return g.val }
func (g *IntGauge) Labels() map[string]string { return g.labels }
func (g *IntGauge) MetricType() MetricType    { return MetricTypeGauge }

func (g *IntGauge) Set(val int) {}
