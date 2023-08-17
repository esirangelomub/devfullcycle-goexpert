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
	t := template.Must(template.New("Curso Template").Parse("Nome: {{.Nome}}\nCarga Horária: {{.CargaHoraria}}h"))
	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
