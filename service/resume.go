package service

import (
	"project-portofolio/model"
	"project-portofolio/repository"

	"go.uber.org/zap"
)

type ResumeService struct {
	Repo   repository.ResumeRepositoryInterface
	Logger *zap.Logger
}

type ResumeServiceInterface interface {
	GetResume() ([]model.Resume, error)
}

// constructor
func NewResumeService(repo repository.ResumeRepositoryInterface, logger *zap.Logger) ResumeService {
	return ResumeService{
		Repo:   repo,
		Logger: logger,
	}
}

func (s *ResumeService) GetResume() ([]model.Resume, error) {
	Resumes, err := s.Repo.Resume()
	if err != nil {
		s.Logger.Error("error  GetResume service ", zap.Error(err))
		return nil, err
	}
	return Resumes, nil
}
