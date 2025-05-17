package web

import "github.com/jgr142/portfolio/internal/models"

type templateData struct {
	Project  models.Project
	Projects []models.Project
}
