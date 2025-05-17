package web

import (
	"log/slog"
	"net/http"
	"strconv"
)

type Handler struct {
	dal    DAL
	logger *slog.Logger
}

func InitHandlers(dal DAL, logger *slog.Logger) Handler {
	return Handler{dal: dal, logger: logger}
}

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	projects, err := h.dal.Latest()
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	data := templateData{Projects: projects}

	h.render(w, r, files, data)
}

func (h *Handler) projectView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		h.clientError(w, http.StatusNotFound)
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/view.tmpl.html",
	}

	project, err := h.dal.Get(id)
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	data := templateData{Project: project}

	h.render(w, r, files, data)
}

func (h *Handler) projectCreate(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/create.tmpl.html",
	}

	h.render(w, r, files, templateData{})
}

func (h *Handler) projectCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Save a new project..."))
}
