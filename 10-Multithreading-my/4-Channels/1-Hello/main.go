package main

import "fmt"

func main() {
	c := make(chan string) // channel empty
	go func() {
		c <- "Hello World" // channel filled
	}()

	msg := <-c // channel empties
	fmt.Println(msg)
}
