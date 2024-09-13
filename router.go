package main

import (
	"io/fs"
	"net/http"

	"github.com/nurbu5/web-template/internal/controllers"
	"github.com/nurbu5/web-template/internal/views"
)

// configRoutes sets up the routing for the application.
// It also parses the layout files sets up the handlers.
func configRoutes(templateFS fs.FS) *http.ServeMux {
	t := views.ConfigViews(templateFS)
	mux := http.NewServeMux()

	mux.HandleFunc("/", controllers.Root(t, templateFS))
	mux.HandleFunc("/test", controllers.Test(t, templateFS))

	return mux
}
