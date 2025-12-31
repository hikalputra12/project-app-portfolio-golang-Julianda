package service

import (
	"project-portofolio/model"
	"project-portofolio/repository"

	"go.uber.org/zap"
)

type SkillService struct {
	Repo   repository.SkillRepositoryInterface
	Logger *zap.Logger
}

type SkillServiceInterface interface {
	GetSkill() ([]model.Skill, error)
}

// constructor
func NewSkillService(repo repository.SkillRepositoryInterface, logger *zap.Logger) SkillService {
	return SkillService{
		Repo:   repo,
		Logger: logger,
	}
}

func (s *SkillService) GetSkill() ([]model.Skill, error) {
	skills, err := s.Repo.Skill()
	if err != nil {
		s.Logger.Error("error  GetSkill service ", zap.Error(err))
		return nil, err
	}
	return skills, nil
}
