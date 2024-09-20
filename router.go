package main

import (
	"net/http"

	"github.com/nurbu5/web-template/internal/controllers"
)

// configRoutes sets up the routing for the application.
// It also parses the layout files sets up the handlers.
func configRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	configStaticFiles(mux)

	mux.HandleFunc("/", controllers.Root())
	mux.HandleFunc("/test", controllers.Test())

	return mux
}
