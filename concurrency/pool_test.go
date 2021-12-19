package concurrency

import (
	"fmt"
	"testing"
)

func TestPool(t *testing.T) {
	concurrency := 5
	pool := NewSimplePool(concurrency)
	jobs := make([]Job, 0)
	jobChan := make(chan int)
	executor := func(num int) int {
		fmt.Printf("Hello!! I'm function %d!!\n", num)
		return num
	}
	for i := 0; i < concurrency + 5; i++ {

		jobs = append(jobs, Job{
			number: i,
			execute: executor,
			result: jobChan,
		})
	}

	for _, job := range jobs {
		pool.Submit(job)
	}

	for x := 0; x < concurrency + 5; x++{
		select {
		case i := <-jobChan :
			fmt.Printf("Recieved response from %d\n", i)
		}
	}

	pool.Stop()
}


