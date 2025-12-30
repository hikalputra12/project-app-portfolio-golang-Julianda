package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"project-portofolio/service"
)

type ResumeHandler struct {
	ServiceResume service.ResumeServiceInterface
	Templates     *template.Template
}

func NewResumeHandler(serviceResume service.ResumeServiceInterface, template *template.Template) ResumeHandler {
	return ResumeHandler{
		ServiceResume: serviceResume,
		Templates:     template,
	}
}

func (h *ResumeHandler) GetResume(w http.ResponseWriter, r *http.Request) {
	Resume, err := h.ServiceResume.GetResume()
	if err != nil {
		fmt.Println("Error dari servoce:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.Templates.ExecuteTemplate(w, "resume", Resume); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
