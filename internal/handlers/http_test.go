package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/go_chantest/internal/core/service"
	"example.com/go_chantest/internal/repositories"
)

var (
	httpHandler = NewHttpHandlerTest()
)

func TestGETIndex(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	httpHandler.Index(response, request)

	t.Run("Get Index", func(t *testing.T){
		if response.Code != http.StatusOK {
			t.Errorf("got HTTP status code %d, expected %d", response.Code, http.StatusOK)
		}
	})
}

func NewHttpHandlerTest() *HTTPHandler {
	websiteRepository := repositories.LoadMemKVS()
	websiteService := service.New(websiteRepository)
	websiteHandler := NewHttpHandler(websiteService)

	return websiteHandler
}
