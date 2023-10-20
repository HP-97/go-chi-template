package handlers

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"example.com/go_chantest/internal/core/ports"
	"github.com/go-chi/chi"
)

//go:embed templates/*
var content embed.FS

//go:embed public
var publicEmbed embed.FS
var publicContent fs.FS

var indexTmpl *template.Template

func init() {
	data, err := content.ReadFile("templates/index.html")
	if err != nil {
		log.Fatalf("%s", err)
	}

	indexTmpl, err = template.New("index").Parse(string(data))
	if err != nil {
		log.Fatalf("%s", err)
	}

	publicFS := fs.FS(publicEmbed)
	publicContent, err = fs.Sub(publicFS, "public")
	if err != nil {
		log.Fatal(err)
	}
}

type HTTPHandler struct {
	websiteService ports.WebsiteService
}

func NewHttpHandler(websiteService ports.WebsiteService) *HTTPHandler {
	return &HTTPHandler{
		websiteService: websiteService,
	}
}

func (hdl *HTTPHandler) Index(w http.ResponseWriter, r *http.Request) {
	indexTmpl.Execute(w, nil)
}

func (hdl *HTTPHandler) Public() fs.FS {
	return publicContent
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(http.FS(publicContent)))
		fs.ServeHTTP(w, r)
	})
}
