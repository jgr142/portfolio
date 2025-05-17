package db

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jgr142/portfolio/internal/models"
)

type ProjectModel struct {
	DB *sql.DB
}

func Open(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func New(db *sql.DB) *ProjectModel {
	return &ProjectModel{DB: db}
}

func (m *ProjectModel) Get(id int) (models.Project, error) {
	stmt := `SELECT id, title, description, img, created FROM projects
	WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	var project models.Project

	err := row.Scan(
		&project.ID,
		&project.Title,
		&project.Description,
		&project.Image,
		&project.Created,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Project{}, ErrNoRecord
		} else {
			return models.Project{}, err
		}
	}

	return project, nil
}

func (m *ProjectModel) Latest() ([]models.Project, error) {
	stmt := `SELECT id, title, description, img, created FROM projects
	ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return make([]models.Project, 0), err
	}

	defer rows.Close()

	var projects []models.Project

	for rows.Next() {
		var project models.Project
		err = rows.Scan(
			&project.ID,
			&project.Title,
			&project.Description,
			&project.Image,
			&project.Created,
		)
		if err != nil {
			return make([]models.Project, 0), err
		}

		projects = append(projects, project)
	}
	if err = rows.Err(); err != nil {
		return make([]models.Project, 0), err
	}

	return projects, nil
}

func (m *ProjectModel) Insert(
	title string,
	desc string,
	img string,
	created int,
) (int, error) {
	stmt := `INSERT INTO snippets (title, description, img, created)
	VALUES(?, ?, ?, ?,))`

	res, err := m.DB.Exec(stmt, title, desc, img, created)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return int(id), err
	}

	return int(id), nil
}
