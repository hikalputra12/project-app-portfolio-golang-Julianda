package handler

import (
	"html/template"
	"project-portofolio/service"
)

type Handler struct {
	BiodataHandler    BiodataHandler
	AboutHandler      AboutHandler
	ResumeHandler     ResumeHandler
	PortofolioHandler PortofolioHandler
	MessageHandler    MessageHandler
}

func NewHandler(service service.Service, template *template.Template) Handler {
	return Handler{
		BiodataHandler:    NewBiodataHandler(&service.Biodata, template),
		AboutHandler:      NewAboutHandler(&service.About, &service.Skill, template),
		ResumeHandler:     NewResumeHandler(&service.Resume, template),
		PortofolioHandler: NewPortofolioHandler(&service.Portofolio, template),
		MessageHandler:    NewMessageHandler(&service.Message, template),
	}
}
