package handlers

import (
	"embed"
	"log"
	"net/http"

	"example.com/go_chantest/internal/core/ports"
)

//go:embed templates/*
var content embed.FS

type HTTPHandler struct {
	websiteService ports.WebsiteService
}

func NewHttpHandler(websiteService ports.WebsiteService) *HTTPHandler {
	return &HTTPHandler{
		websiteService: websiteService,
	}
}

func (hdl *HTTPHandler) Index(w http.ResponseWriter, r *http.Request) {
	data, err := content.ReadFile("templates/index.html")
	if err != nil {
		log.Printf("%s", err)
	}
	w.Write([]byte(data))
}

