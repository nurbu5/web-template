package controllers

import "net/http"

// Path: /
func Root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
