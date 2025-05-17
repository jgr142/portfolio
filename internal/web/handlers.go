package web

import (
	"log/slog"
	"net/http"
	"strconv"
)

type Handler struct {
	dal    DAL
	tCache templateCache
	logger *slog.Logger
}

func InitHandlers(dal DAL, tCache templateCache, logger *slog.Logger) Handler {
	return Handler{dal: dal, tCache: tCache, logger: logger}
}

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	projects, err := h.dal.Latest()
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	data := NewTemplateData()
	data.Projects = projects

	h.render(w, r, "home", data)
}

func (h *Handler) projectView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		h.clientError(w, http.StatusNotFound)
		return
	}

	project, err := h.dal.Get(id)
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	data := NewTemplateData()
	data.Project = project

	h.render(w, r, "view", data)
}

func (h *Handler) projectCreate(w http.ResponseWriter, r *http.Request) {
	h.render(w, r, "create", templateData{})
}

func (h *Handler) projectCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Save a new project..."))
}
