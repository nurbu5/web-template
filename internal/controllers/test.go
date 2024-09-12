package controllers

import (
	"html/template"
	"net/http"
)

type testData struct {
	Test string
}

// Path: /test
func Test(layouts *template.Template) http.HandlerFunc {
	c := func(layouts *template.Template) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			layouts.Execute(w, testData{"TESTER'S"})
		}
	}

	return newController(
		layouts.Lookup("base.html"),
		"Root",
		[]string{"test.html"},
		c,
	).handler
}
