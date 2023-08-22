package main

import (
	"fmt"
	"github.com/esirangelomub/devfullcycle-goexpert/7-Packaging/1-my/math"
)

func main() {
	m := math.NewMath(3, 4)
	m.C = 12
	fmt.Println(m)
	fmt.Println(m.Add())
	fmt.Println(m.C)
	fmt.Println(math.X)
}
