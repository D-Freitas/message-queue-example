package main

import "golang-message-queue/worker"

func main() {
	defer worker.Wait()
	for i := 0; i < 20; i++ {
		job := worker.Job{
			Action:  PrintPayload,
			Payload: i,
		}
		job.Fire()
	}
}
