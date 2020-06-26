package monitoring

import (
	"context"
)

type Writer interface {
	Write(ctx context.Context, events Events) error
}

type DummyWriter struct{}

func NewDummyWriter(url string) (Writer, error) {
	return &DummyWriter{}, nil
}

func (w *DummyWriter) Write(ctx context.Context, events Events) error {
	return nil
}
