package gauge

type Gauge interface {
	Set(val interface{}) error
}
