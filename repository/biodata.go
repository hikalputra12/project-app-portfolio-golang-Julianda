package repository

import (
	"context"
	"project-portofolio/database"
	"project-portofolio/model"
)

type BiodataRepository struct {
	db database.PgxIface
}
type BiodataRepositoryInterface interface {
	Biodata() (*model.Biodata, error)
}

// constructor
func NewBiodataRepository(db database.PgxIface) BiodataRepository {
	return BiodataRepository{
		db: db,
	}
}

func (r *BiodataRepository) Biodata() (*model.Biodata, error) {
	query := `SELECT full_name, biodata FROM biodata;`

	var s model.Biodata
	err := r.db.QueryRow(context.Background(), query).Scan(
		&s.Name,
		&s.Biodata,
	)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
