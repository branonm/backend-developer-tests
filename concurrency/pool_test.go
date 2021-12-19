package concurrency

import (
	"fmt"
	"testing"
)

func TestPool(t *testing.T) {
	concurrency := 15
	jobCount := 10
	pool := NewSimplePool(concurrency)
	jobs := make([]Job, 0)
	jobChan := make(chan int)

	// This is the 'work' that is done
	executor := func(num int) int {
		fmt.Printf("Hello!! I'm function %d!!\n", num)
		return num
	}

	// Allow for jobChan to demux responses from Jobs
	for i := 0; i < jobCount; i++ {

		jobs = append(jobs, Job{
			number:  i,
			execute: executor,
			result:  jobChan,
		})
	}

	for _, job := range jobs {
		pool.Submit(job)
	}

	for x := 0; x < jobCount; x++ {
		select {
		case i := <-jobChan:
			fmt.Printf("Recieved response from %d\n", i)
		}
	}
	pool.Stop()
}
