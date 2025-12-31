package repository

import (
	"context"
	"project-portofolio/database"
	"project-portofolio/model"

	"go.uber.org/zap"
)

type BiodataRepository struct {
	db     database.PgxIface
	Logger *zap.Logger
}
type BiodataRepositoryInterface interface {
	Biodata() (*model.Biodata, error)
}

// constructor
func NewBiodataRepository(db database.PgxIface, logger *zap.Logger) BiodataRepository {
	return BiodataRepository{
		db:     db,
		Logger: logger,
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
		r.Logger.Error("error query biodata repo ", zap.Error(err))
		return nil, err
	}

	return &s, nil
}
