package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"project-portofolio/service"
)

type AboutHandler struct {
	ServiceAbout service.AboutServiceInterface
	Templates    *template.Template
}

func NewABoutHandler(serviceAbout service.AboutServiceInterface, template *template.Template) AboutHandler {
	return AboutHandler{
		ServiceAbout: serviceAbout,
		Templates:    template,
	}
}

func (h *AboutHandler) GetAbout(w http.ResponseWriter, r *http.Request) {
	about, err := h.ServiceAbout.GetAbout()
	if err != nil {
		fmt.Println("Error dari servoce:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//set header
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//eksekusi templatenya
	if err := h.Templates.ExecuteTemplate(w, "about", about); err != nil {
		//cek error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
