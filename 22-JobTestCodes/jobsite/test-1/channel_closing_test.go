package main

import (
	"testing"
	"time"
)

func TestChannelClosedAfterValues(t *testing.T) {
	ch := make(chan int)

	// Sample functions for testing
	fn1 := func() int { time.Sleep(1 * time.Second); return 1 }
	fn2 := func() int { time.Sleep(2 * time.Second); return 2 }
	fn3 := func() int { time.Sleep(3 * time.Second); return 3 }

	go executeParallel(ch, fn1, fn2, fn3)

	// This function reads from the channel and checks if it's closed.
	isChannelClosed := func(ch <-chan int) bool {
		select {
		case _, ok := <-ch:
			if !ok {
				return true // The channel is closed.
			}
		default:
		}
		return false
	}

	// We use a timer to prevent this from running indefinitely.
	timeout := time.After(10 * time.Second)

	// Loop until either the channel is closed or we hit a timeout.
	for {
		select {
		case <-timeout:
			t.Fatal("Test timed out! Channel might not be closed.")
			return
		default:
			if isChannelClosed(ch) {
				return // Test passed.
			}
			time.Sleep(500 * time.Millisecond) // Check every half second.
		}
	}
}
