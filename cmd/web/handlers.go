package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	tp, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	err = tp.Execute(w, nil)
	if err != nil {
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}
}

func (app *application) projectView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/view.tmpl.html",
	}

	tp, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	err = tp.Execute(w, nil)
	if err != nil {
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}
}

func (app *application) projectCreate(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/create.tmpl.html",
	}

	tp, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	err = tp.Execute(w, nil)
	if err != nil {
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Save a new project..."))
}
