package broker

type Producer interface {
	Publish(body []byte) error
	Close() error
}
