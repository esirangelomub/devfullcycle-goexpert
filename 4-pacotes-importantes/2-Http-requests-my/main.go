package main

import (
	"io"
	"net/http"
)

func main() {
	request, err := http.Get("https://www.google.com.br")
	if err != nil {
		panic(err)
	}
	defer request.Body.Close()
	res, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))
	//request.Body.Close()
}
