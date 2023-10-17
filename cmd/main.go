package main

import (
	"log"
	"net/http"

	"example.com/go_chantest/internal/config"
	"example.com/go_chantest/internal/core/service"
	"example.com/go_chantest/internal/handlers"
	"example.com/go_chantest/internal/repositories"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	config.InitConfig()

	websiteRepository := repositories.LoadMemKVS()
	websiteService := service.New(websiteRepository)
	websiteHandler := handlers.NewHttpHandler(websiteService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", websiteHandler.Index)

	log.Printf("running server on :3000!")
	http.ListenAndServe(":3000", r)
}
