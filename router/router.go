package router

import (
	"net/http"
	"project-portofolio/handler"
	"project-portofolio/middleware"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func NewRouter(h handler.Handler, log *zap.Logger) *chi.Mux {
	r := chi.NewRouter()
	//penggunaan logger
	r.Use(middleware.Logging(log))

	r.Group(func(r chi.Router) {
		r.Get("/", h.BiodataHandler.GetBiodata)
		r.Get("/about", h.AboutHandler.GetAboutPage)
		r.Get("/portofolio", h.PortofolioHandler.GetPortofolio)
		r.Get("/resume", h.ResumeHandler.GetResume)
		r.Get("/contact", h.MessageHandler.GetContactPage)
		r.Post("/message", h.MessageHandler.PostMessage)

	})

	fs := http.FileServer(http.Dir("assets"))
	r.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	return r
}
