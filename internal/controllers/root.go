package controllers

import (
	"html/template"
	"io/fs"
	"net/http"
)

// Path: /
func Root(layouts *template.Template, templateFS fs.FS) http.HandlerFunc {
	c := func(layouts *template.Template) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			layouts.Execute(w, nil)
		}
	}

	return newController(
		layouts.Lookup("base.html"), // The layout template
		"Root",                      // The name of the controller
		[]string{"root.html"},       // The template files to parse for this request
		c,                           // The controller function that handles the request logic
		templateFS,                  // The template file system
	).handler
}
