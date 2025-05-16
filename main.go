package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Joshua!"))
}

func projectView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific project..."))
}

func projectCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new project..."))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/project/view/{id}", projectView)
	mux.HandleFunc("/project/create", projectCreate)
	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
