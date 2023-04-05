package worker

import (
	"sync"
)

var wg sync.WaitGroup

type Job struct {
	Action  func(int)
	Payload int
}

func Wait() {
	wg.Wait()
}

func (job Job) Fire() {
	defer wg.Add(1)
	go func() {
		job.Action(job.Payload)
		wg.Done()
	}()
}
