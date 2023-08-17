package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	t := template.Must(template.ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Cursos{
		Curso{"Go", 40},
		Curso{"Java", 60},
		Curso{"Python", 50},
	})
	if err != nil {
		panic(err)
	}
}
