package repository

import (
	"context"
	"project-portofolio/database"
	"project-portofolio/model"
)

type SkillRepository struct {
	db database.PgxIface
}
type SkillRepositoryInterface interface {
	Skill() ([]model.Skill, error)
}

// constructor
func NewSkillRepository(db database.PgxIface) SkillRepository {
	return SkillRepository{
		db: db,
	}
}

func (r *SkillRepository) Skill() ([]model.Skill, error) {
	query := `SELECT  name FROM skills`

	var skill []model.Skill
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var s model.Skill
		err := rows.Scan(
			&s.Name)
		if err != nil {
			return nil, err
		}
		skill = append(skill, s)
	}

	return skill, nil
}
