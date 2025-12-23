package service

import (
	"project-portofolio/model"
	"project-portofolio/repository"
)

type BiodataService struct {
	Repo repository.BiodataRepositoryInterface
}

type BiodataServiceInterface interface {
	GetBiodata() (*model.Biodata, error)
}

// constructor
func NewBiodataService(repo repository.BiodataRepositoryInterface) BiodataService {
	return BiodataService{
		Repo: repo,
	}
}

func (s *BiodataService) GetBiodata() (*model.Biodata, error) {
	biodata, err := s.Repo.Biodata()
	if err != nil {
		return nil, err
	}
	return biodata, nil
}
