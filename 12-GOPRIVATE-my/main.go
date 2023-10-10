package main

import (
	"fmt"
	"github.com/esirangelomub/fcutils-secret/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
