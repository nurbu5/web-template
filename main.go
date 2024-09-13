package main

import (
	"embed"
	"net/http"
)

//go:embed web/templates
var templateFS embed.FS

func main() {
	mux := configRoutes(templateFS)

	http.ListenAndServe(":8080", mux)
}
