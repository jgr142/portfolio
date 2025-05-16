package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /project/view/{id}", app.projectView)
	mux.HandleFunc("GET /project/create", app.projectCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	return mux
}
