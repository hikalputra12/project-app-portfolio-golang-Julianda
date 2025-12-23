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
	Education() ([]model.Education, error)
	Experience() ([]model.Education, error)
}

// constructor
func NewResumeRepository(db database.PgxIface) ResumeRepository {
	return ResumeRepository{
		db: db,
	}
}

func (r *ResumeRepository) Education() ([]model.Education, error) {
	query := `SELECT institution, degree, major, start_year, end_year, description FROM education`

	var education []model.Education
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var s model.Education
		err := rows.Scan(
			&s.Institution,
			&s.Degree,
			&s.Major,
			&s.Start_year,
			&s.End_year,
			&s.Description,
		)
		if err != nil {
			return nil, err
		}
		education = append(education, s)
	}

	return education, nil

}
func (r *ResumeRepository) Experience() ([]model.Experience, error) {
	query := `SELECT  title, company, type, start_year, end_year, description  FROM experience`

	var experience []model.Experience
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var s model.Experience
		err := rows.Scan(
			&s.Title,
			&s.Company,
			&s.Type,
			&s.Start_year,
			&s.End_year,
			&s.Description,
		)
		if err != nil {
			return nil, err
		}
		experience = append(experience, s)
	}

	return experience, nil

}

func (r *ResumeRepository) Skill() ([]model.Skill, error) {
	query := `SELECT  name, category, level, icon FROM skills`

	var skill []model.Skill
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var s model.Skill
		err := rows.Scan(
			&s.Name,
			&s.Category,
			&s.Level,
			&s.Icon,
		)
		if err != nil {
			return nil, err
		}
		skill = append(skill, s)
	}

	return skill, nil
}
