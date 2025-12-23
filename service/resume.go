package service

import (
	"project-portofolio/model"
	"project-portofolio/repository"
)

type ResumeService struct {
	Repo repository.ResumeRepositoryInterface
}

type ResumeServiceInterface interface {
	GetEducation() ([]model.Education, error)
	GetExperience() ([]model.Experience, error)
	GetSkill() ([]model.Skill, error)
}

// constructor
func NewResumeService(repo repository.ResumeRepositoryInterface) ResumeService {
	return ResumeService{
		Repo: repo,
	}
}

func (s *ResumeService) GetEducation() ([]model.Education, error) {
	educations, err := s.Repo.Education()
	if err != nil {
		return nil, err
	}
	return educations, nil
}

func (s *ResumeService) GetExperience() ([]model.Experience, error) {
	experiences, err := s.Repo.Experience()
	if err != nil {
		return nil, err
	}
	return experiences, nil
}

func (s *ResumeService) GetSkill() ([]model.Skill, error) {
	skills, err := s.Repo.Skill()
	if err != nil {
		return nil, err
	}
	return skills, nil
}
