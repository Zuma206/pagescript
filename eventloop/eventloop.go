package eventloop

type Eventloop struct {
	continuations chan *Continuation
	tasks         chan *Task
}

func (eventloop *Eventloop) Go(function TaskFunction) {
	eventloop.tasks <- &Task{function}
}

func (eventloop *Eventloop) Continue(function ContinuationFunction) *Continuation {
	continuation := NewContinuation(function)
	eventloop.continuations <- continuation
	return continuation
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
