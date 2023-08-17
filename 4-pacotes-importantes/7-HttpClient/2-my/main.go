package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	json := bytes.NewBuffer([]byte(`{"name":"Eduardo"}`))
	resp, err := c.Post("https://google.com", "application/json", json)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
