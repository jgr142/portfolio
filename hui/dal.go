package web

import (
	"github.com/jgr142/portfolio/internal/models"
)

type DAL interface {
	Get(id int) (models.Project, error)
	Latest() ([]models.Project, error)
	Insert(title string, description string, img string, created int) (int, error)
}
