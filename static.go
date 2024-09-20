package main

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed internal/views/web/assets/dist
var staticFS embed.FS

func configStaticFiles(mux *http.ServeMux) {
	// Strip the "internal/views/web/assets/dist" prefix to make the embedded files accessible at "/static"
	staticFS, err := fs.Sub(staticFS, "internal/views/web/assets/dist")
	if err != nil {
		panic(err)
	}

	// Serve static files from embedded filesystem
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))
}
