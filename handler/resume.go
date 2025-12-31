package handler

import (
	"html/template"
	"net/http"
	"project-portofolio/service"

	"go.uber.org/zap"
)

type ResumeHandler struct {
	ServiceResume service.ResumeServiceInterface
	Templates     *template.Template
	Logger        *zap.Logger
}

func NewResumeHandler(serviceResume service.ResumeServiceInterface, template *template.Template, logger *zap.Logger) ResumeHandler {
	return ResumeHandler{
		ServiceResume: serviceResume,
		Templates:     template,
		Logger:        logger,
	}
}

func (h *ResumeHandler) GetResume(w http.ResponseWriter, r *http.Request) {
	Resume, err := h.ServiceResume.GetResume()
	if err != nil {
		h.Logger.Error("error  GetResume handler ", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.Templates.ExecuteTemplate(w, "resume", Resume); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
