package web

import (
	"html/template"
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

	tp, err := template.ParseFiles(files...)
	if err != nil {
		h.logger.Error("internal server error", slog.String("err", err.Error()))
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	data, err := h.dal.Latest()
	if err != nil {
		h.logger.Error("internal server error", slog.String("err", err.Error()))
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	err = tp.Execute(w, data)
	if err != nil {
		h.logger.Error("internal server error", slog.String("err", err.Error()))
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}
}

func (h *Handler) projectView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		h.logger.Error("client error", slog.String("err", err.Error()))
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/view.tmpl.html",
	}

	tp, err := template.ParseFiles(files...)
	if err != nil {
		h.logger.Error("internal server error", slog.String("err", err.Error()))
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	data, err := h.dal.Get(id)
	if err != nil {
		h.logger.Error("internal server error", slog.String("err", err.Error()))
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	err = tp.Execute(w, data)
	if err != nil {
		h.logger.Error("internal server error", slog.String("err", err.Error()))
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}
}

func (h *Handler) projectCreate(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/create.tmpl.html",
	}

	tp, err := template.ParseFiles(files...)
	if err != nil {
		h.logger.Error("internal server error", slog.String("err", err.Error()))
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	err = tp.Execute(w, nil)
	if err != nil {
		h.logger.Error("internal server error", slog.String("err", err.Error()))
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}
}

func (h *Handler) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Save a new project..."))
}
