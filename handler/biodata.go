package handler

import (
	"html/template"
	"net/http"
	"project-portofolio/service"

	"go.uber.org/zap"
)

type BiodataHandler struct {
	ServiceBiodata service.BiodataServiceInterface
	Templates      *template.Template
	Logger         *zap.Logger
}

func NewBiodataHandler(serviceBiodata service.BiodataServiceInterface, template *template.Template, logger *zap.Logger) BiodataHandler {
	return BiodataHandler{
		ServiceBiodata: serviceBiodata,
		Templates:      template,
		Logger:         logger,
	}
}

func (h *BiodataHandler) GetBiodata(w http.ResponseWriter, r *http.Request) {
	biodata, err := h.ServiceBiodata.GetBiodata()
	if err != nil {
		h.Logger.Error("error  GetBiodata handler ", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.Templates.ExecuteTemplate(w, "home", biodata); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
