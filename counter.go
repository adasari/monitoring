package monitoring

import (
	"sync/atomic"
)

type IntCounter struct {
	name   string
	val    uint64
	labels map[string]string

	m *Monitoring
}

func NewIncCounter(name string, lables map[string]string) *IntCounter {
	return nil
}

func (c *IntCounter) Inc() {
	atomic.AddUint64(&c.val, 1)
}

func (c *IntCounter) Name() string              { return c.name }
func (c *IntCounter) Value() interface{}        { return c.val }
func (c *IntCounter) Labels() map[string]string { return c.labels }
func (c *IntCounter) MetricType() MetricType    { return MetricTypeCounter }
