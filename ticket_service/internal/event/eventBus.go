package event

type ChannelSignalBus struct {
	Channel chan struct{}
}

func NewChannelSignalBus() *ChannelSignalBus {
	return &ChannelSignalBus{
		Channel: make(chan struct{}, 100),
	}
}

func (eb *ChannelSignalBus) Publish() {
	eb.Channel <- struct{}{}
}

func (eb *ChannelSignalBus) Subscribe() <-chan struct{} {
	return eb.Channel
}
