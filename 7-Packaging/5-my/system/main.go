package main

import (
	"github.com/esirangelomub/devfullcycle-goexpert/7-Packaging/5-my/math"
	"github.com/google/uuid"
)

func main() {
	m := math.NewMath(4, 8)
	println(m.Add())
	println(uuid.New().String())
}
