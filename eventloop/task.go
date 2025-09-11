package eventloop

type TaskFunction func() ContinuationFunction
type Task struct {
	function TaskFunction
}
