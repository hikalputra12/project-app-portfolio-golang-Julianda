package repository

import (
	"context"
	"project-portofolio/database"
	"project-portofolio/model"
)

type PortofolioRepository struct {
	db database.PgxIface
}
type PortofolioRepositoryInterface interface {
	Portofolio() ([]model.Portofolio, error)
}

// constructor
func NewPortofolioRepository(db database.PgxIface) PortofolioRepository {
	return PortofolioRepository{
		db: db,
	}
}

func (r *PortofolioRepository) Portofolio() ([]model.Portofolio, error) {
	query := `SELECT 	title, 
	slug,
	description,
	thumbnail,
	tech_stack,
	github_url FROM projects`

	var Portofolio []model.Portofolio
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var s model.Portofolio
		err := rows.Scan(
			&s.Title,
			&s.Slug,
			&s.Description,
			&s.Github_url,
			&s.Tech_stack,
			&s.Github_url,
		)
		if err != nil {
			return nil, err
		}
		Portofolio = append(Portofolio, s)
	}

	return Portofolio, nil

}
