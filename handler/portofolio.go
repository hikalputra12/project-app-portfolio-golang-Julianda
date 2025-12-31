package handler

import (
	"html/template"
	"net/http"
	"project-portofolio/service"

	"go.uber.org/zap"
)

type PortofolioHandler struct {
	ServicePortofolio service.PortofolioServiceInterface
	Templates         *template.Template
	Logger            *zap.Logger
}

func NewPortofolioHandler(servicePortofolio service.PortofolioServiceInterface, template *template.Template, logger *zap.Logger) PortofolioHandler {
	return PortofolioHandler{
		ServicePortofolio: servicePortofolio,
		Templates:         template,
		Logger:            logger,
	}
}

func (h *PortofolioHandler) GetPortofolio(w http.ResponseWriter, r *http.Request) {
	portofolio, err := h.ServicePortofolio.GetPortofolio()
	if err != nil {
		h.Logger.Error("error  GetPortofolio handler ", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.Templates.ExecuteTemplate(w, "portofolio", portofolio); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
