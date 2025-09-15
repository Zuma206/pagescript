package eventloop

import "github.com/Zuma206/pagescript/options"

const (
	defaultContinuationQueueSize = 32
	defaultTaskQueueSize         = 32
)

func NewEventloop(opts ...options.Option[*Eventloop]) *Eventloop {
	eventloop := &Eventloop{
		continuations: make(chan Continuation, defaultContinuationQueueSize),
		tasks:         make(chan *Task, defaultTaskQueueSize),
		err:           make(chan error, 1),
	}
	if err := options.Apply(eventloop, opts); err != nil {
		return nil
	}
	return eventloop
}

var WithContinuationQueueSize = options.New(func(eventloop *Eventloop, size uint) error {
	eventloop.continuations = make(chan Continuation, size)
	return nil
})

var WithTaskQueueSize = options.New(func(eventloop *Eventloop, size uint) error {
	eventloop.tasks = make(chan *Task, size)
	return nil
})

var WithInitialWorkers = options.New(func(eventloop *Eventloop, n uint) error {
	eventloop.Workers(n)
	return nil
})
