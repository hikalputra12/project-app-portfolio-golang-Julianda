package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"project-portofolio/database"
	"project-portofolio/handler"
	"project-portofolio/repository"
	"project-portofolio/router"
	"project-portofolio/service"
	"project-portofolio/utils"
)

func main() {
	var templates = template.Must(template.New("").ParseGlob("templates/*.html"))
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}
	for _, t := range templates.Templates() {
		fmt.Println("template:", t.Name())
	}

	logger, err := utils.InitLogger("./logs/app-", true)
	repo := repository.NewRepository(db, logger)
	service := service.NewService(repo, logger)
	handler := handler.NewHandler(service, templates, logger)
	r := router.NewRouter(handler, logger)

	fmt.Println("server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("error server")
	}

}
