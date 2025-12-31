package repository

import (
	"context"
	"project-portofolio/database"
	"project-portofolio/model"

	"go.uber.org/zap"
)

type AboutRepository struct {
	db     database.PgxIface
	Logger *zap.Logger
}
type AboutRepositoryInterface interface {
	About() (*model.About, error)
}

// constructor
func NewAboutRepository(db database.PgxIface, logger *zap.Logger) AboutRepository {
	return AboutRepository{
		db:     db,
		Logger: logger,
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
		r.Logger.Error("error query about repo ", zap.Error(err))
		return nil, err
	}

	return &s, nil
}
