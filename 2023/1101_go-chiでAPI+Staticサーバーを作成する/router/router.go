package router

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"time"

	"github.com/a10adotapp/blog.a10a.app/2023/1025/router/helper"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func NewRouter(fsys fs.FS) *chi.Mux {
	tmpl, err := template.ParseFS(fsys, []string{
		"static/index.html",
	}...)
	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()

	router.Use(middleware.Timeout(10 * time.Second))
	router.Use(middleware.Logger)

	router.Route("/api", func(r chi.Router) {
		r.Get("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			fmt.Fprint(w, "{\"message\":\"hello from api\"}")
		}))
	})

	router.Route("/", func(r chi.Router) {
		r.Get("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			pageData := &struct {
				Message string
			}{
				Message: "hello from server",
			}

			if err := tmpl.ExecuteTemplate(w, helper.ParseTemplateFilename(r.URL.Path), pageData); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
		}))
	})

	return router
}
