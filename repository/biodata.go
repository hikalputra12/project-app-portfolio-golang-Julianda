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
	About() (*model.Biodata, error)
}

// constructor
func NewBiodataRepository(db database.PgxIface) BiodataRepository {
	return BiodataRepository{
		db: db,
	}
}

func (r *BiodataRepository) Biodata() (*model.Biodata, error) {
	query := `SELECT full_name,title,address,phone,email FROM biodata;`

	var s model.Biodata
	err := r.db.QueryRow(context.Background(), query).Scan(
		&s.Name,
		&s.Address,
		&s.Profession,
		&s.Email,
		&s.Phone,
	)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
func (r *BiodataRepository) About() (*model.About, error) {
	query := `SELECT about FROM biodata;`

	var s model.About
	err := r.db.QueryRow(context.Background(), query).Scan(
		&s.About,
	)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
