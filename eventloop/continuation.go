package eventloop

type ContinuationFunction func() error
type Continuation struct {
	function ContinuationFunction
	channel  chan error
}

func NewContinuation(function ContinuationFunction) *Continuation {
	return &Continuation{
		channel:  make(chan error, 1),
		function: function,
	}
}

func (continuation *Continuation) Await() error {
	return <-continuation.channel
}

func (continuation *Continuation) Run() error {
	err := continuation.function()
	continuation.channel <- err
	return err
}
