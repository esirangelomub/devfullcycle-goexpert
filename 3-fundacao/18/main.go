package main

import (
	"fmt"
	"goexpert/3-fundacao/18/matematica"
)

func main() {
	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}

	fmt.Println(carro.Andar())
	fmt.Println("Resultado: ", s)
	fmt.Println(matematica.A)
}
