package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/nurbu5/web-template/internal/controllers"
)

//go:embed internal/views/web/assets/dist
var staticFS embed.FS

// configRoutes sets up the routing for the application.
// It also parses the layout files sets up the handlers.
func configRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	configStaticFiles(mux)

	mux.HandleFunc("/", controllers.Root())
	mux.HandleFunc("/test", controllers.Test())

	return mux
}

func configStaticFiles(mux *http.ServeMux) {
	// Strip the "internal/views/web/assets/dist" prefix to make the embedded files accessible at "/static"
	staticFS, err := fs.Sub(staticFS, "internal/views/web/assets/dist")
	if err != nil {
		panic(err)
	}

	// Serve static files from embedded filesystem
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))
}
