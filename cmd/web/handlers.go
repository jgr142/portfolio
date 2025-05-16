package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Joshua!"))
}

func projectView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display a specific project with ID %d...", id)
	w.Write([]byte(msg))
}

func projectCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new project..."))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Save a new project..."))
}
