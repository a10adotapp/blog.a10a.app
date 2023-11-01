package main

import (
	"embed"
	"net/http"

	"github.com/a10adotapp/blog.a10a.app/2023/1025/router"
	"github.com/joho/godotenv"
)

//go:embed static/*
var staticFS embed.FS

func main() {
	godotenv.Load()

	http.ListenAndServe(":3000", router.NewRouter(staticFS))
}
