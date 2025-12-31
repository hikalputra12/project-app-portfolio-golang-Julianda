package service

import (
	"project-portofolio/repository"

	"go.uber.org/zap"
)

type Service struct {
	Biodata    BiodataService
	About      AboutService
	Resume     ResumeService
	Portofolio PortofolioService
	Message    MessageService
	Skill      SkillService
}

func NewService(repo repository.Repository, logger *zap.Logger) Service {
	return Service{
		Biodata:    NewBiodataService(&repo.Biodata, logger),
		About:      NewAboutService(&repo.About, logger),
		Resume:     NewResumeService(&repo.Resume, logger),
		Portofolio: NewPortofolioService(&repo.Portofolio, logger),
		Message:    NewMessageService(&repo.Message, logger),
		Skill:      NewSkillService(&repo.Skill, logger),
	}
}
