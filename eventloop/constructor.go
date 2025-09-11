package eventloop

type EventloopOptions struct {
	continuationQueueSize uint
	taskQueueSize         uint
}

type EventloopOption func(options *EventloopOptions)

func newEventloopOptions(options []EventloopOption) *EventloopOptions {
	optionsStruct := &EventloopOptions{
		continuationQueueSize: 32,
		taskQueueSize:         32,
	}
	for _, option := range options {
		option(optionsStruct)
	}
	return optionsStruct
}

func NewEventloop(options ...EventloopOption) *Eventloop {
	opts := newEventloopOptions(options)
	return &Eventloop{
		continuationQueue: make(chan *Continuation, opts.continuationQueueSize),
		taskQueue:         make(chan *Task, opts.continuationQueueSize),
	}
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
