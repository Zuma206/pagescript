package eventloop

import (
	"sync"
	"testing"
	"time"
)

const (
	simulatedWorkLength = 10 * time.Millisecond
	nContinuations      = 10
)

func TestContinuationsRunInSerial(t *testing.T) {
	var mu sync.Mutex
	var wg sync.WaitGroup
	loop := NewEventloop(WithContinuationQueueSize(nContinuations))

	for i := range nContinuations {
		wg.Add(1)
		loop.Continue(func() error {
			if !mu.TryLock() {
				t.Errorf("continuation %d ran in parallel", i)
			} else {
				time.Sleep(simulatedWorkLength)
				mu.Unlock()
			}
			wg.Done()
			return nil
		})
	}

	go loop.Start()
	wg.Wait()
	loop.Stop()
}

const (
	nTasks = 10
)

func TestTasksRunInParallel(t *testing.T) {
	loop := NewEventloop(WithTaskQueueSize(nTasks))
	expectedTime := nTasks * simulatedWorkLength
	startTime := time.Now()
	var wg sync.WaitGroup

	for range nTasks {
		wg.Add(1)
		loop.Go(func() ContinuationFunction {
			time.Sleep(simulatedWorkLength)
			wg.Done()
			return nil
		})
	}

	go loop.Worker()
	go loop.Worker()
	wg.Wait()
	loop.Stop()

	tasksTime := time.Since(startTime)
	if tasksTime >= expectedTime {
		t.Errorf(
			"tasks took %dms, which is more than/equal to the estimated serialised time (%dms)",
			tasksTime.Milliseconds(),
			expectedTime.Milliseconds(),
		)
	}
}
