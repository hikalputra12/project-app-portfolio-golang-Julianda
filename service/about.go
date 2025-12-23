package service

import (
	"project-portofolio/model"
	"project-portofolio/repository"
)

type AboutService struct {
	Repo repository.AboutRepositoryInterface
}

type AboutServiceInterface interface {
	GetAbout() (*model.About, error)
}

// constructor
func NewAboutService(repo repository.AboutRepositoryInterface) AboutService {
	return AboutService{
		Repo: repo,
	}
}

func (s *AboutService) GetAbout() (*model.About, error) {
	about, err := s.Repo.About()
	if err != nil {
		return nil, err
	}
	return about, nil
}
