package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/go_chantest/internal/config"
	"example.com/go_chantest/internal/core/service"
	"example.com/go_chantest/internal/handlers"
	"example.com/go_chantest/internal/repositories"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	filepathPtr := flag.String("path", "", "path to website config file")
	flag.Parse()

	config.InitConfig()

	f, err := os.ReadFile(*filepathPtr)
	if err != nil {
		log.Fatal(fmt.Sprintf("unable to print file %s", *filepathPtr))
	}
	fStr := string(f)

	fmt.Printf("file %s contains %s\n", *filepathPtr, fStr)

	websiteRepository := repositories.LoadMemKVS(fStr)
	websiteService := service.New(websiteRepository)
	websiteHandler := handlers.NewHttpHandler(websiteService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", websiteHandler.GetAll)

	http.ListenAndServe(":3000", r)
}
