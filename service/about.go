package service

import (
	"project-portofolio/model"
	"project-portofolio/repository"

	"go.uber.org/zap"
)

type AboutService struct {
	Repo   repository.AboutRepositoryInterface
	Logger *zap.Logger
}

type AboutServiceInterface interface {
	GetAbout() (*model.About, error)
}

// constructor
func NewAboutService(repo repository.AboutRepositoryInterface, logger *zap.Logger) AboutService {
	return AboutService{
		Repo:   repo,
		Logger: logger,
	}
}

func (s *AboutService) GetAbout() (*model.About, error) {
	about, err := s.Repo.About()
	if err != nil {
		s.Logger.Error("error  GetAbout service ", zap.Error(err))
		return nil, err
	}
	return about, nil
}
