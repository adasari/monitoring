package monitoring

import (
	"atomic/sync"
	"fmt"
	"log"
	"time"
)

// EventHandler is interface that provides abstract functions for the metric envent processing.
type EventHandler interface {
	handle(event Events)
}

type Stopper interface {
	Stop() error
}

type Collector interface {
	Collect()
}

// eventHandler is buffered event handler.
type eventHandler struct {
	ch   chan Events
	done chan struct{}

	mu     sync.Mutex
	events Events

	flushThreshold int
	flushTicker    *time.Ticker
}

// NewEventHandler returns buffered event handler instance for given flushThreshold and flushInterval.
func NewEventHandler(flushThreshold int, flushInterval time.Duration) (EventHandler, error) {
	ch := make(chan Events)
	done := make(chan struct{})
	// flushThreshold or flushInterval is mandatory
	if flushInterval <= 0 && flushThreshold <= 0 {
		return nil, fmt.Errorf("invalid flushThreshold and flushInterval: %v, %v", flushThreshold, flushInterval)
	}

	eh := &eventHandler{
		ch:             ch,
		done:           done,
		flushThreshold: flushThreshold,
	}

	if flushThreshold < 0 {
		eh.flushThreshold = -1
	}

	if flushInterval > 0 {
		eh.flushTicker = time.NewTicker(flushInterval)

		go func() {
			for {
				select {
				case <-eh.done:
					log.Printf("monitoring is stopped")
					eh.stop()
					return
				case <-eh.flushTicker.C:
					eh.flush()
				}
			}
		}()
	}

	return eh, nil
}

// handle is an implementation of EventHandler handle function. it buffers events and flush based on event handler flushThreshold.
func (eh *eventHandler) handle(events Events) {
	eh.mu.Lock()
	defer eh.mu.Unlock()

	for _, e := range events {
		eh.events = append(eh.events, e)
		if len(eh.events) >= eh.flushThreshold {
			eh.flushUnlocked()
		}
	}
}

func (eh *eventHandler) stop() error {
	eh.flush()
	eh.flushTicker.Stop()
	return nil
}

// flush pushes the events from events slice to event channel with lock.
func (eh *eventHandler) flush() {
	eh.mu.Lock()
	defer eh.mu.Unlock()

	eh.flushUnlocked()
}

// flushUnlocked pushes the events from events slice to event channel without lock.
func (eh *eventHandler) flushUnlocked() {
	if len(eh.events) > 0 {
		eh.ch <- eh.events
		eh.events = []Event{}
	}
}

// unBufferedEventHandler is buffered event handler.
type unBufferedEventHandler struct {
	ch   chan Events
	done chan struct{}
}

// NewUnBufferedEventHandler .
func NewUnBufferedEventHandler(ch chan Events, done chan struct{}) (EventHandler, error) {
	ueh := &unBufferedEventHandler{
		ch:   ch,
		done: done,
	}

	return ueh, nil
}

// handle is an implementation of EventHandler handle function. it buffers events and flush based on event handler flushThreshold.
func (ueh *unBufferedEventHandler) handle(events Events) {
	ueh.ch <- events
}

func (ueh *unBufferedEventHandler) stop() error {
	return nil
}
