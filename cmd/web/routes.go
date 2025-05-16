package main

import "net/http"

func routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /project/view/{id}", projectView)
	mux.HandleFunc("GET /project/create", projectCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	return mux
}
