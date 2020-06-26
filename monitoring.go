package monitoring

import (
	"context"
	"log"
	"time"
)

type Monitoring struct {
	ch chan Events
	w  Writer

	eventHandler EventHandler
}

// New .
func New(w Writer, flushInterval, flushThreshold int) (*Monitoring, error) {
	if flushInterval <= 0 && flushThreshold <= 0 {
		// TODO: return error
	}

	eh, err := NewEventHandler(flushThreshold, time.Duration(flushInterval)*time.Second)
	if err != nil {
		return nil, err
	}

	return &Monitoring{
		w:            w,
		eventHandler: eh,
	}, nil
}

// Collect collects events from channel and pass them to writer. it is applicable for only async event handlers.
func (m *Monitoring) collect() {
	for {
		events, ok := <-m.ch
		if !ok {
			log.Printf("event channel is closed")
			return
		}

		if err := m.w.Write(context.Background(), events); err != nil {
			log.Printf("failed to write metric events %+v: %v", events, err)
		}
	}

}

func (m *Monitoring) Close() error {
	if m.eventHandler != nil {
		m.eventHandler.stop()
	}
}

// NewIntCounter .
func (m *Monitoring) NewIntCounter(projectID string, name string, labels map[string]string) (*IntCounter, error) {
	// TODO: any validations required? Add it.
	c := &IntCounter{
		name:   name,
		labels: labels,
		m:      m,
	}

	return c, nil
}
