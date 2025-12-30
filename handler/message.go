package handler

import (
	"html/template"
	"net/http"
	"project-portofolio/model"
	"project-portofolio/service"
)

type MessageHandler struct {
	MessageService service.MessageServiceInterface
	Template       *template.Template
}

func NewMessageHandler(messageService service.MessageServiceInterface, template *template.Template) MessageHandler {
	return MessageHandler{
		MessageService: messageService,
		Template:       template,
	}
}

// untuk merender halaman message
func (h *MessageHandler) GetContactPage(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Contact Me",
	}

	// 2. Set Header
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := h.Template.ExecuteTemplate(w, "contact", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *MessageHandler) PostMessage(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	formMessage := r.FormValue("message")

	if name == "" || email == "" || formMessage == "" {
		data := map[string]interface{}{
			"Title": "Contact Me",
			"Error": "Semua kolom wajib diisi!",
		}
		h.Template.ExecuteTemplate(w, "contact", data)
		return
	}

	newMessage := model.Message{
		FullName: name,
		Email:    email,
		Message:  formMessage,
	}

	err := h.MessageService.CreateNewMessage(&newMessage)

	if err != nil {
		http.Error(w, "Gagal menyimpan pesan: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	data := map[string]interface{}{
		"Title":   "Contact Me",
		"Success": true, // Ini akan memicu alert hijau di HTML
	}

	if err := h.Template.ExecuteTemplate(w, "contact", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
