package repository

import (
	"context"
	"project-portofolio/database"
	"project-portofolio/model"

	"go.uber.org/zap"
)

type ResumeRepository struct {
	db     database.PgxIface
	Logger *zap.Logger
}
type ResumeRepositoryInterface interface {
	Resume() ([]model.Resume, error)
}

// constructor
func NewResumeRepository(db database.PgxIface, logger *zap.Logger) ResumeRepository {
	return ResumeRepository{
		db:     db,
		Logger: logger,
	}
}

func (r *ResumeRepository) Resume() ([]model.Resume, error) {
	query := `SELECT title, organization, period, description FROM resume`

	var Resume []model.Resume
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		r.Logger.Error("error query message repo ", zap.Error(err))
		return nil, err
	}

	for rows.Next() {
		var s model.Resume
		err := rows.Scan(
			&s.Title,
			&s.Organization,
			&s.Periode,
			&s.Description,
		)
		if err != nil {
			return nil, err
		}
		Resume = append(Resume, s)
	}

	return Resume, nil

}
