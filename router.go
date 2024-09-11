package main

import "net/http"

func configRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)

	return mux
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
