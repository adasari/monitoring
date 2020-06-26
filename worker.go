package monitoring

type Worker interface {
	EventHandler
	Stopper
	Collector
}
