package handlers

import (
	"fmt"
	"net/http"

	"example.com/go_chantest/internal/core/ports"
)

type HTTPHandler struct {
	websiteService ports.WebsiteService
}

func NewHttpHandler(websiteService ports.WebsiteService) *HTTPHandler {
	return &HTTPHandler{
		websiteService: websiteService,
	}
}

func (hdl *HTTPHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	websites := hdl.websiteService.GetAll()
	w.Write([]byte(fmt.Sprintf("websites = %v", websites)))
}

