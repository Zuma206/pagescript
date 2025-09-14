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
