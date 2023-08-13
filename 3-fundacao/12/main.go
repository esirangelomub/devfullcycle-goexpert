package main

func main() {
	// Memória -> Endreço -> Valor
	a := 10
	var ponteiro *int = &a
	*ponteiro = 20
	b := &a
	*b = 30
	println(ponteiro)
	println(*ponteiro)
	println(&a)
	println(a)
	println(b)
	println(*b)
}
