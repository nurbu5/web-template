package main

import (
	"net/http"
)

func main() {
	mux := configRoutes()

	http.ListenAndServe(":8080", mux)
}
