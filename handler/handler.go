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
}

func NewHandler(service service.Service, template *template.Template) Handler {
	return Handler{
		BiodataHandler:    NewBiodataHandler(&service.Biodata, template),
		AboutHandler:      NewABoutHandler(&service.About, template),
		ResumeHandler:     NewResumeHandler(&service.Resume, template),
		PortofolioHandler: NewPortofolioHandler(&service.Portofolio, template),
	}
}
