package service

import (
	"project-portofolio/model"
	"project-portofolio/repository"
)

type PortofolioService struct {
	Repo repository.PortofolioRepositoryInterface
}

type PortofolioServiceInterface interface {
	GetPortofolio() ([]model.Portofolio, error)
}

// constructor
func NewPortofolioService(repo repository.PortofolioRepositoryInterface) PortofolioService {
	return PortofolioService{
		Repo: repo,
	}
}

func (s *PortofolioService) GetPortofolio() ([]model.Portofolio, error) {
	Portofolio, err := s.Repo.Portofolio()
	if err != nil {
		return nil, err
	}
	return Portofolio, nil
}
