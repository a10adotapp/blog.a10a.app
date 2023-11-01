package main

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/a10adotapp/blog.a10a.app/2023/1025/router"
	"github.com/joho/godotenv"
)

//go:embed static/*
var staticFS embed.FS

func main() {
	godotenv.Load()

	tmpl, err := template.ParseFS(staticFS, []string{
		"static/index.html",
	}...)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":3000", router.NewRouter(tmpl))
}
