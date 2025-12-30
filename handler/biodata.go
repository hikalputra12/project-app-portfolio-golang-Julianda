package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"project-portofolio/service"
)

type BiodataHandler struct {
	ServiceBiodata service.BiodataServiceInterface
	Templates      *template.Template
}

func NewBiodataHandler(serviceBiodata service.BiodataServiceInterface, template *template.Template) BiodataHandler {
	return BiodataHandler{
		ServiceBiodata: serviceBiodata,
		Templates:      template,
	}
}

func (h *BiodataHandler) GetBiodata(w http.ResponseWriter, r *http.Request) {
	biodata, err := h.ServiceBiodata.GetBiodata()
	if err != nil {
		fmt.Println("Error dari servoce:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//set header
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//eksekusi templatenya
	if err := h.Templates.ExecuteTemplate(w, "home", biodata); err != nil {
		//cek error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
