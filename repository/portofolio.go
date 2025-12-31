package repository

import (
	"context"
	"project-portofolio/database"
	"project-portofolio/model"

	"go.uber.org/zap"
)

type PortofolioRepository struct {
	db     database.PgxIface
	Logger *zap.Logger
}
type PortofolioRepositoryInterface interface {
	Portofolio() ([]model.Portofolio, error)
}

// constructor
func NewPortofolioRepository(db database.PgxIface, logger *zap.Logger) PortofolioRepository {
	return PortofolioRepository{
		db:     db,
		Logger: logger,
	}
}

func (r *PortofolioRepository) Portofolio() ([]model.Portofolio, error) {
	query := `SELECT title, description, tech_stack, github_url FROM projects`

	var Portofolio []model.Portofolio
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		r.Logger.Error("error query message repo ", zap.Error(err))
		return nil, err
	}

	for rows.Next() {
		var s model.Portofolio
		err := rows.Scan(
			&s.Title,
			&s.Description,
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
