package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"project-portofolio/model"
	"project-portofolio/service"
)

type AboutHandler struct {
	ServiceAbout service.AboutServiceInterface
	ServiceSkill service.SkillServiceInterface // <-- Tambahan PENTING
	Templates    *template.Template
}

func NewAboutHandler(serviceAbout service.AboutServiceInterface,
	serviceSkill service.SkillServiceInterface, template *template.Template) AboutHandler {
	return AboutHandler{
		ServiceAbout: serviceAbout,
		ServiceSkill: serviceSkill,
		Templates:    template,
	}
}

func (h *AboutHandler) GetAboutPage(w http.ResponseWriter, r *http.Request) {
	aboutResult, err := h.ServiceAbout.GetAbout()
	if err != nil {
		fmt.Println("Error dari service:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	skillResult, err := h.ServiceSkill.GetSkill()
	if err != nil {
		fmt.Println("Error dari service:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data := model.AboutPageData{
		IntroData: aboutResult,
		SkillData: skillResult,
	}

	// 4. Masak (Render)
	h.Templates.ExecuteTemplate(w, "about", data)
}
