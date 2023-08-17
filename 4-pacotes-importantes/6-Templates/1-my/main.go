package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"Go", 40}
	tmp := template.New("Curso Template")
	tmp, _ = tmp.Parse("Nome: {{.Nome}}\nCarga Horária: {{.CargaHoraria}}h")
	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
