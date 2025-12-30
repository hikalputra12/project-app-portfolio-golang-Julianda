package service

import (
	"project-portofolio/model"
	"project-portofolio/repository"
)

type ResumeService struct {
	Repo repository.ResumeRepositoryInterface
}

type ResumeServiceInterface interface {
	GetResume() ([]model.Resume, error)
}

// constructor
func NewResumeService(repo repository.ResumeRepositoryInterface) ResumeService {
	return ResumeService{
		Repo: repo,
	}
}

func (s *ResumeService) GetResume() ([]model.Resume, error) {
	Resumes, err := s.Repo.Resume()
	if err != nil {
		return nil, err
	}
	return Resumes, nil
}
