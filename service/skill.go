package service

import (
	"project-portofolio/model"
	"project-portofolio/repository"
)

type SkillService struct {
	Repo repository.SkillRepositoryInterface
}

type SkillServiceInterface interface {
	GetSkill() ([]model.Skill, error)
}

// constructor
func NewSkillService(repo repository.SkillRepositoryInterface) SkillService {
	return SkillService{
		Repo: repo,
	}
}

func (s *SkillService) GetSkill() ([]model.Skill, error) {
	skills, err := s.Repo.Skill()
	if err != nil {
		return nil, err
	}
	return skills, nil
}
