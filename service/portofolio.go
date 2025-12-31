package service

import (
	"project-portofolio/model"
	"project-portofolio/repository"

	"go.uber.org/zap"
)

type PortofolioService struct {
	Repo   repository.PortofolioRepositoryInterface
	Logger *zap.Logger
}

type PortofolioServiceInterface interface {
	GetPortofolio() ([]model.Portofolio, error)
}

// constructor
func NewPortofolioService(repo repository.PortofolioRepositoryInterface, logger *zap.Logger) PortofolioService {
	return PortofolioService{
		Repo:   repo,
		Logger: logger,
	}
}

func (s *PortofolioService) GetPortofolio() ([]model.Portofolio, error) {
	Portofolio, err := s.Repo.Portofolio()
	if err != nil {
		s.Logger.Error("error  GetPortofolio service ", zap.Error(err))
		return nil, err
	}
	return Portofolio, nil
}
