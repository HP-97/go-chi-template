package main

import (
	"fmt"
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

	c := config.GetConfig()

	websiteRepository := repositories.LoadMemKVS()
	websiteService := service.New(websiteRepository)
	websiteHandler := handlers.NewHttpHandler(websiteService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", websiteHandler.Index)

	serverStr := fmt.Sprintf("%s:%s", c.GetString("host.addr"),c.GetString("host.port"))
	log.Printf("running server on %s!", serverStr)
	http.ListenAndServe(serverStr, r)
}
