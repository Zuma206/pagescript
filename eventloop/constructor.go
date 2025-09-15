package eventloop

type EventloopOptions struct {
	continuationQueueSize uint
	taskQueueSize         uint
	initialWorkers        uint
}

type EventloopOption func(options *EventloopOptions)

func newEventloopOptions(options []EventloopOption) *EventloopOptions {
	optionsStruct := &EventloopOptions{
		continuationQueueSize: 32,
		taskQueueSize:         32,
		initialWorkers:        0,
	}
	for _, option := range options {
		option(optionsStruct)
	}
	return optionsStruct
}

func NewEventloop(options ...EventloopOption) *Eventloop {
	opts := newEventloopOptions(options)
	eventloop := &Eventloop{
		continuations: make(chan Continuation, opts.continuationQueueSize),
		tasks:         make(chan *Task, opts.continuationQueueSize),
	}
	eventloop.Workers(opts.initialWorkers)
	return eventloop
}

func WithContinuationQueueSize(continuationQueueSize uint) EventloopOption {
	return func(options *EventloopOptions) {
		options.continuationQueueSize = continuationQueueSize
	}
}

func WithTaskQueueSize(taskQueueSize uint) EventloopOption {
	return func(options *EventloopOptions) {
		options.taskQueueSize = taskQueueSize
	}
}

func WithInitialWorkers(initialWorkers uint) EventloopOption {
	return func(options *EventloopOptions) {
		options.initialWorkers = initialWorkers
	}
}
