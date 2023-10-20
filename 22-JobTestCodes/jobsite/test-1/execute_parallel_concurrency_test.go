package main

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestExecuteParallelConcurrency(t *testing.T) {
	// Define two functions that both sleep for a set duration
	shortDuration := 100 * time.Millisecond
	longDuration := 200 * time.Millisecond

	shortFunction := func() int {
		time.Sleep(shortDuration)
		return 1
	}
	longFunction := func() int {
		time.Sleep(longDuration)
		return 2
	}

	ch := make(chan int, 2) // buffer the channel so it doesn't block

	// Start timing
	startTime := time.Now()

	executeParallel(ch, shortFunction, longFunction)

	// Wait for the results
	for i := 0; i < 2; i++ {
		<-ch
	}

	// End timing
	endTime := time.Now()

	totalDuration := endTime.Sub(startTime)

	// If the total duration is significantly greater than longDuration, the functions probably didn't run in parallel
	if totalDuration > longDuration+10*time.Millisecond { // adding a small buffer time
		t.Errorf("Functions don't seem to be running in parallel. Expected duration ~%v, got %v", longDuration, totalDuration)
	}
}

func TestExecuteParallelConcurrencyMultiple(t *testing.T) {
	// This counter will track the number of functions running concurrently.
	var concurrentCounter int32

	// Define a mock function that simulates a function taking time to execute.
	mockFunction := func(duration time.Duration) func() int {
		return func() int {
			atomic.AddInt32(&concurrentCounter, 1)
			time.Sleep(duration)
			atomic.AddInt32(&concurrentCounter, -1)
			return 1
		}
	}

	shortFunction := mockFunction(1 * time.Second)
	mediumFunction := mockFunction(2 * time.Second)
	longFunction := mockFunction(3 * time.Second)

	ch := make(chan int)

	// Start the functions in parallel.
	go executeParallel(ch, shortFunction, mediumFunction, longFunction)

	// After a short time, at least two functions should be running concurrently.
	time.Sleep(500 * time.Millisecond)
	if count := atomic.LoadInt32(&concurrentCounter); count < 2 {
		t.Fatalf("Expected at least 2 functions to run concurrently, but got %d", count)
	}

	// Drain the channel.
	for _ = range ch {
	}
}
