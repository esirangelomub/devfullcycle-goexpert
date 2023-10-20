package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestExecuteParallel(t *testing.T) {
	// Define the input functions for the test
	expensiveFunction := func() int { return exampleFunction(200000000) }
	cheapFunction := func() int { return exampleFunction(10000000) }

	// The expected results (the order might vary since goroutines are concurrent)
	expected := []int{10000000, 200000000}

	ch := make(chan int)
	go executeParallel(ch, expensiveFunction, cheapFunction)

	// Read the results from the channel
	var results []int
	for val := range ch {
		results = append(results, val)
	}

	// Sort the results to ensure consistent comparison with expected values
	sort.Ints(results)

	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}
