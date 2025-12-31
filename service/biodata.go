package service

import (
	"project-portofolio/model"
	"project-portofolio/repository"

	"go.uber.org/zap"
)

type BiodataService struct {
	Repo   repository.BiodataRepositoryInterface
	Logger *zap.Logger
}

type BiodataServiceInterface interface {
	GetBiodata() (*model.Biodata, error)
}

// constructor
func NewBiodataService(repo repository.BiodataRepositoryInterface, logger *zap.Logger) BiodataService {
	return BiodataService{
		Repo:   repo,
		Logger: logger,
	}
}

func (s *BiodataService) GetBiodata() (*model.Biodata, error) {
	biodata, err := s.Repo.Biodata()
	if err != nil {
		s.Logger.Error("error  GetBiodata service ", zap.Error(err))
		return nil, err
	}
	return biodata, nil
}
