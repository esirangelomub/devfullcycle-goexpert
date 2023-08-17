package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	//mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("Olá mundo"))
	//})
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "My Blog"})
	http.ListenAndServe("8080", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Olá Eduardo"))
	})
	http.ListenAndServe("8081", mux2)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo"))
}

type blog struct {
	title string
}

// Para que uma interface (struct) tenha methods, esta é a forma de implementar,
// usando func (b blog) ...
func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
