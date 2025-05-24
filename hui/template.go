package web

import (
	"bytes"
	"errors"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/jgr142/portfolio/internal/models"
)

type templateData struct {
	currentYear int
	Project     models.Project
	Projects    []models.Project
}

func NewTemplateData() templateData {
	return templateData{currentYear: time.Now().Year()}
}

type templateCache struct {
	cache map[string]*template.Template
}

func (h *Handler) render(
	w http.ResponseWriter,
	r *http.Request,
	pageName string,
	data templateData,
) {
	tp, err := h.tCache.Get(pageName)
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	buf := new(bytes.Buffer)
	err = tp.Execute(buf, data)
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	buf.WriteTo(w)
}

func NewTemplateCache() (templateCache, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return templateCache{}, err
	}

	for _, page := range pages {
		name := strings.TrimSuffix(filepath.Base(page), ".tmpl.html")

		tp, err := template.ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return templateCache{}, err
		}

		tp, err = tp.ParseFiles(page)
		if err != nil {
			return templateCache{}, err
		}

		cache[name] = tp
	}

	return templateCache{cache: cache}, nil
}

func (tc *templateCache) Get(pageName string) (*template.Template, error) {
	tp, ok := tc.cache[pageName]
	if !ok {
		return nil, errors.New("template does not exist")
	}

	return tp, nil
}
