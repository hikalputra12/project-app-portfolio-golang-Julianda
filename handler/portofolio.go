package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"project-portofolio/service"
)

type PortofolioHandler struct {
	ServicePortofolio service.PortofolioServiceInterface
	Templates         *template.Template
}

func NewPortofolioHandler(servicePortofolio service.PortofolioServiceInterface, template *template.Template) PortofolioHandler {
	return PortofolioHandler{
		ServicePortofolio: servicePortofolio,
		Templates:         template,
	}
}

func (h *PortofolioHandler) GetPortofolio(w http.ResponseWriter, r *http.Request) {
	portofolio, err := h.ServicePortofolio.GetPortofolio()
	if err != nil {
		fmt.Println("Error dari servoce:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.Templates.ExecuteTemplate(w, "portofolio", portofolio); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
