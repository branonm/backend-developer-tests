// Package concurrency implements worker pool interfaces, one simple and one a
// bit more complex.
package concurrency

import (
	"fmt"
	"time"
)

type Job struct {
	number int
	execute func(int) int
	result chan int
}

// SimplePool is a simple worker pool that does not support cancellation or
// closing. All functions are safe to call from multiple goroutines.
type SimplePool interface {
	// Submit a task to be executed asynchronously. This function will return as
	// soon as the task is submitted. If the pool does not have an available slot
	// for the task, this blocks until it can submit.
	Submit(Job)
	Stop()
}

func worker(jobs <-chan Job) {

	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				fmt.Println("Stopping worker")
				return
			}
			job.result <-job.execute(job.number)
		default:
			fmt.Println("Waiting for Job.")
			time.Sleep(1 * time.Second)

		}
	}
}

type Pool struct {
	maxConcurrency int
	tasks          chan Job
}

// NewSimplePool creates a new SimplePool that only allows the given maximum
// concurrent tasks to run at any one time. maxConcurrent must be greater than
// zero.
func NewSimplePool(maxConcurrent int) SimplePool {

	pool := &Pool{
		maxConcurrency: maxConcurrent,
		tasks:          make(chan Job, maxConcurrent),
	}

	pool.run()

	return pool
}

func (p *Pool) Submit(job Job) {
	p.tasks <- job
}

func (p *Pool) run() {

	for i := 0; i < p.maxConcurrency; i++ {
		go worker(p.tasks)
	}
}

func (p *Pool) Stop() {
	close(p.tasks)

	// allow worker routines to
	time.Sleep(10 * time.Second)
}
