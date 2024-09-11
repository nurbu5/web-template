package main

import (
	"net/http"

	"github.com/nurbu5/web-template/internal/controllers"
)

func configRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controllers.Root)

	return mux
}
