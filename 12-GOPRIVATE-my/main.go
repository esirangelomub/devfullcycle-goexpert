package main

import (
	"fmt"
	"github.com/esirangelomub/fcutils-secret/pkg/events"
)

/*
Some commands to run:
- go mod init github.com/esirangelomub/fcutils-secret
- go mod tidy
Environment configuration:
- .netrc file in the home directory
  - github.com	<username>	<token>
  - api.bitbucket.org	<username>	<token>

- export GOPRIVATE="github.com/esirangelomub/*"
*/
func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
