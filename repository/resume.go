package repository

import (
	"context"
	"project-portofolio/database"
	"project-portofolio/model"
)

type ResumeRepository struct {
	db database.PgxIface
}
type ResumeRepositoryInterface interface {
	Resume() ([]model.Resume, error)
}

// constructor
func NewResumeRepository(db database.PgxIface) ResumeRepository {
	return ResumeRepository{
		db: db,
	}
}

func (r *ResumeRepository) Resume() ([]model.Resume, error) {
	query := `SELECT title, organization, period, description FROM resume`

	var Resume []model.Resume
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
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
