package handler

import (
	"html/template"
	"project-portofolio/service"

	"go.uber.org/zap"
)

type Handler struct {
	BiodataHandler    BiodataHandler
	AboutHandler      AboutHandler
	ResumeHandler     ResumeHandler
	PortofolioHandler PortofolioHandler
	MessageHandler    MessageHandler
}

func NewHandler(service service.Service, template *template.Template, logger *zap.Logger) Handler {
	return Handler{
		BiodataHandler:    NewBiodataHandler(&service.Biodata, template, logger),
		AboutHandler:      NewAboutHandler(&service.About, &service.Skill, template, logger),
		ResumeHandler:     NewResumeHandler(&service.Resume, template, logger),
		PortofolioHandler: NewPortofolioHandler(&service.Portofolio, template, logger),
		MessageHandler:    NewMessageHandler(&service.Message, template, logger),
	}
}
