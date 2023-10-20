package main

import (
	"fmt"
	"sync"
)

//• Functions received as the second argument should be executed in parallel.
//• As soon as a function finishes, its result should be written to the channel ch.
//• After all functions finish, channel ch should be closed, and executeParallel should finish its execution.
//
//For example, executing main function should produce the following console output:
//Result: 10000000
//Result: 200000000
//
//It's not working:
//Example case: Wrong answer
//Functions are executed in parallel: Wrong answer
//Channel returns correct values: Correct answer
//Channel is closed after it returns correct values: Wrong answer

// executeParallel runs multiple functions in parallel and sends their results to a channel.
//
// Parameters:
// - ch: The channel where function results are sent.
// - functions: A variadic parameter that accepts an arbitrary number of functions. Each function should return an int.
func executeParallel(ch chan<- int, functions ...func() int) {
	// A WaitGroup waits for a collection of goroutines to finish.
	// The main goroutine calls Add to set the number of goroutines to wait for.
	var wg sync.WaitGroup
	wg.Add(len(functions))

	// Loop over each function
	for _, fn := range functions {
		// Launch the function inside a goroutine.
		go func(f func() int) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Execute the function and send its result to the channel.
			ch <- f()
		}(fn)
	}

	// Another goroutine that waits for all functions to finish, then closes the channel.
	go func() {
		// Wait blocks until the WaitGroup counter is zero.
		wg.Wait()
		// Close the channel after all the functions have executed.
		close(ch)
	}()
}

// exampleFunction returns a sum after iterating for a given number of times.
//
// Parameters:
// - counter: The number of iterations.
func exampleFunction(counter int) int {
	sum := 0
	for i := 0; i < counter; i++ {
		sum += 1
	}
	return sum
}

func main() {
	// Define two functions for testing purposes.

	// expensiveFunction represents a longer-running task.
	expensiveFunction := func() int {
		fmt.Println("Expensive Function Started")
		// This function runs for a longer time because it has more iterations.
		result := exampleFunction(200000000)
		fmt.Println("Expensive Function Ended")
		return result
	}

	// cheapFunction represents a shorter-running task.
	cheapFunction := func() int {
		fmt.Println("Cheap Function Started")
		// This function runs for a shorter time due to fewer iterations.
		result := exampleFunction(10000000)
		fmt.Println("Cheap Function Ended")
		return result
	}

	// Create a channel to receive results from the functions.
	ch := make(chan int)

	// Run the functions in parallel.
	executeParallel(ch, expensiveFunction, cheapFunction)

	// Read and print results from the channel until it's closed.
	for result := range ch {
		fmt.Printf("Result: %d\n", result)
	}
}
