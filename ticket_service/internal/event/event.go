package event

type EventSignalBus interface {
	Publish()
	Subscribe() <-chan struct{}
}
