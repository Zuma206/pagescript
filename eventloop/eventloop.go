package eventloop

type Eventloop struct {
	continuations chan Continuation
	tasks         chan *Task
}

func (eventloop *Eventloop) Go(function TaskFunction) {
	eventloop.tasks <- &Task{function}
}

func (eventloop *Eventloop) Callback(function ContinuationFunction) {
	callback := &Callback{function}
	eventloop.continuations <- callback
}

func (eventloop *Eventloop) Block(function ContinuationFunction) error {
	block := &Block{
		err:      make(chan error),
		function: function,
	}
	eventloop.continuations <- block
	return <-block.err
}

func (eventloop *Eventloop) Start() error {
	for continuation := range eventloop.continuations {
		if err := continuation.Run(); err != nil {
			return err
		}
	}
	return nil
}

func (eventloop *Eventloop) Stop() {
	close(eventloop.tasks)
	close(eventloop.continuations)
}

func (eventloop *Eventloop) Worker() {
	for task := range eventloop.tasks {
		if continuation := task.function(); continuation != nil {
			eventloop.Callback(continuation)
		}
	}
}

func (eventloop *Eventloop) Workers(n uint) {
	for range n {
		go eventloop.Worker()
	}
}
