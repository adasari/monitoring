package monitoring

// Noop does nothing.
type noopWorker struct {
	done chan struct{}
}

func NewNoopWorker() (Worker, error) {
	return &noopWorker{}, nil
}

// handle is an implementation of EventHandler handle function.
func (nh *noopWorker) handle(events Events) {}

func (nh *noopWorker) Stop() error {
	return nil
}

func (nh *noopWorker) Collect() {}
