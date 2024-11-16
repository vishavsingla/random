package worker

import "myproject/job"

const MaxWorkers = 5

type WorkerPool struct {
	JobQueue chan int
}

func NewWorkerPool() *WorkerPool {
	return &WorkerPool{JobQueue: make(chan int, 100)}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < MaxWorkers; i++ {
		go wp.worker()
	}
}

func (wp *WorkerPool) worker() {
	for jobID := range wp.JobQueue {
		job.ProcessJob(jobID)
	}
}

func (wp *WorkerPool) Enqueue(jobID int) {
	wp.JobQueue <- jobID
}
