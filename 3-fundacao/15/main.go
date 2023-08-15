package main

import "fmt"

// Utilizar pouco, pois o ideal é tipar as variáveis
type xis interface{}

func main() {
	var x xis = 10
	var y interface{} = "Hello, World!"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é %T e o valor é %v\n", t, t)
}
