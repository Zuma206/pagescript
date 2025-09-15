package eventloop

type ContinuationFunction func() error
type Continuation interface {
	Run() error
}

type Callback struct {
	function ContinuationFunction
}

func (callback *Callback) Run() error {
	return callback.function()
}

type Block struct {
	function ContinuationFunction
	err      chan error
}

func (block *Block) Run() error {
	err := block.function()
	block.err <- err
	return nil
}
