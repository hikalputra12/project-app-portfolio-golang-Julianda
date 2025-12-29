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

func (h *ResumeHandler) GetEducation(w http.ResponseWriter, r *http.Request) {
	Resume, err := h.ServiceResume.GetEducation()
	if err != nil {
		fmt.Println("Error dari servoce:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//set header
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//eksekusi templatenya
	if err := h.Templates.ExecuteTemplate(w, "Resume", Resume); err != nil {
		//cek error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ResumeHandler) GetExperience(w http.ResponseWriter, r *http.Request) {
	Resume, err := h.ServiceResume.GetExperience()
	if err != nil {
		fmt.Println("Error dari servoce:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//set header
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//eksekusi templatenya
	if err := h.Templates.ExecuteTemplate(w, "Resume", Resume); err != nil {
		//cek error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ResumeHandler) GetSkill(w http.ResponseWriter, r *http.Request) {
	Resume, err := h.ServiceResume.GetSkill()
	if err != nil {
		fmt.Println("Error dari servoce:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//set header
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//eksekusi templatenya
	if err := h.Templates.ExecuteTemplate(w, "Resume", Resume); err != nil {
		//cek error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
