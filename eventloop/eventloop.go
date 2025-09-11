package eventloop

type Eventloop struct {
	continuations chan *Continuation
	tasks         chan *Task
}
