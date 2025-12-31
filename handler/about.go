package handler

import (
	"html/template"
	"net/http"
	"project-portofolio/model"
	"project-portofolio/service"

	"go.uber.org/zap"
)

type AboutHandler struct {
	ServiceAbout service.AboutServiceInterface
	ServiceSkill service.SkillServiceInterface // <-- Tambahan PENTING
	Templates    *template.Template
	Logger       *zap.Logger
}

func NewAboutHandler(serviceAbout service.AboutServiceInterface,
	serviceSkill service.SkillServiceInterface, template *template.Template, logger *zap.Logger) AboutHandler {
	return AboutHandler{
		ServiceAbout: serviceAbout,
		ServiceSkill: serviceSkill,
		Templates:    template,
		Logger:       logger,
	}
}

func (h *AboutHandler) GetAboutPage(w http.ResponseWriter, r *http.Request) {
	aboutResult, err := h.ServiceAbout.GetAbout()
	if err != nil {
		h.Logger.Error("error  GetAboutPage handler ", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	skillResult, err := h.ServiceSkill.GetSkill()
	if err != nil {
		h.Logger.Error("error  GetSkillhandler ", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data := model.AboutPageData{
		IntroData: aboutResult,
		SkillData: skillResult,
	}

	// 4. Masak (Render)
	h.Templates.ExecuteTemplate(w, "about", data)
}
