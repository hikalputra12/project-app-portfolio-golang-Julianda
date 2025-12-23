package service

import (
	"project-portofolio/model"
	"project-portofolio/repository"
)

type BiodataService struct {
	Repo repository.BiodataRepositoryInterface
}

type BiodataServiceInterface interface {
	Biodata() (*model.Biodata, error)
}

// constructor
func NewBiodataRepository(repo repository.BiodataRepositoryInterface) BiodataService {
	return BiodataService{
		Repo: repo,
	}
}

func (s *BiodataService) Biodata() (*model.Biodata, error) {
	biodata, err := s.Repo.Biodata()
	if err != nil {
		return nil, err
	}
	return biodata, nil
}
