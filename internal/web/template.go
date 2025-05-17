package web

import (
	"html/template"
	"net/http"

	"github.com/jgr142/portfolio/internal/models"
)

type templateData struct {
	Project  models.Project
	Projects []models.Project
}

func (h *Handler) render(
	w http.ResponseWriter,
	r *http.Request,
	files []string,
	data templateData,
) {
	tp, err := template.ParseFiles(files...)
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	err = tp.Execute(w, data)
	if err != nil {
		h.serverError(w, r, err)
		return
	}
}
