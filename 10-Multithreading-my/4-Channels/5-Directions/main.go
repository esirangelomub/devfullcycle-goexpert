package main

import "fmt"

// When the arrow is on the right side of the function signature,
// it means that the channel only receives values
// receive only channel
func write(name string, hello chan<- string) {
	hello <- name
}

// When the arrow is on the left side of the function signature,
// it means that the channel only returns (empties the channel) values
// send only channel
func read(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	hello := make(chan string)
	go write("Golang", hello)
	read(hello)
}
