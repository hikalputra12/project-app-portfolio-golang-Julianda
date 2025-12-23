package service

import "project-portofolio/repository"

type Service struct {
	Biodata    BiodataService
	About      AboutService
	Resume     ResumeService
	Portofolio PortofolioService
	Message    MessageService
}

func NewReository(repo repository.Repository) Service {
	return Service{
		Biodata:    NewBiodataService(&repo.Biodata),
		About:      NewAboutService(&repo.About),
		Resume:     NewResumeService(&repo.Resume),
		Portofolio: NewPortofolioService(&repo.Portofolio),
		Message:    NewMessageService(&repo.Message),
	}
}
