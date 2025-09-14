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
