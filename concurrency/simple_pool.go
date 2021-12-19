// Package concurrency implements worker pool interfaces, one simple and one a
// bit more complex.
package concurrency

import (
	"fmt"
	"sync"
)

// Job is a simple abstraction representing a unit of work
type Job struct {
	number  int
	execute func(int) int
	result  chan int
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

// worker executes Jobs passed in via channel
func worker(jobs <-chan Job, wg *sync.WaitGroup, num int) {
	defer wg.Done()
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Stopping worker %d\n", num)
				return
			}
			job.result <- job.execute(job.number)
			//default:
			//	fmt.Println("Waiting for Job.")
		}
	}
}

// Pool is the struct that implements the SimplePool interface
type Pool struct {
	jobs           chan Job
	maxConcurrency int
	wg             sync.WaitGroup
}

// NewSimplePool creates a new SimplePool that only allows the given maximum
// concurrent jobs to run at any one time. maxConcurrent must be greater than
// zero.
func NewSimplePool(maxConcurrent int) SimplePool {

	pool := &Pool{
		maxConcurrency: maxConcurrent,
		jobs:           make(chan Job, maxConcurrent),
	}

	for i := 0; i < pool.maxConcurrency; i++ {
		pool.wg.Add(1)
		fmt.Printf("Starting worker %d\n", i)
		go worker(pool.jobs, &pool.wg, i)
	}

	return pool
}

// Submit pushes a Job on to the Job channel contained in Pool
func (p *Pool) Submit(job Job) {
	p.jobs <- job
}

// Stop closes the Job channel allowing the workers to spin down
func (p *Pool) Stop() {
	close(p.jobs)
	p.wg.Wait()
}
