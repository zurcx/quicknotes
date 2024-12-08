package main

import (
	"fmt"
	"net/http"
)

func noteList(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Listagem de Notas ...")
}

func noteView(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Exibindo um Notas ...")
}

func noteCreate(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Criando uma Nota ...")
}
func main() {
	fmt.Println("teste")

	http.ListenAndServe(":5000", nil)
<<<<<<< HEAD
=======

	mux := http.NewServeMux()

	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/create", noteCreate)

>>>>>>> origin/main
}
