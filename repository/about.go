package repository

import (
	"context"
	"project-portofolio/database"
	"project-portofolio/model"
)

type AboutRepository struct {
	db database.PgxIface
}
type AboutRepositoryInterface interface {
	About() (*model.About, error)
}

// constructor
func NewAboutRepository(db database.PgxIface) AboutRepository {
	return AboutRepository{
		db: db,
	}
}

func (r *AboutRepository) About() (*model.About, error) {
	query := `SELECT intro, focus FROM about;`

	var s model.About
	err := r.db.QueryRow(context.Background(), query).Scan(
		&s.IntroAbout,
		&s.FocusAbout,
	)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
