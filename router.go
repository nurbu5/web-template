package main

import (
	"net/http"

	"github.com/nurbu5/web-template/internal/controllers"
	"github.com/nurbu5/web-template/internal/views"
)

// configRoutes sets up the routing for the application.
// It also parses the layout files sets up the handlers.
func configRoutes() *http.ServeMux {
	t := views.ConfigViews()
	mux := http.NewServeMux()

	mux.HandleFunc("/", controllers.Root(t))
	mux.HandleFunc("/test", controllers.Test(t))

	return mux
}
