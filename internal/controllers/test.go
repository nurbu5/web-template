package controllers

import (
	"net/http"

	"github.com/nurbu5/web-template/internal/views"
)

type testData struct {
	Test string
}

// Path: /test
func Test() http.HandlerFunc {
	c := func(v *views.View) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			v.Execute(w, testData{"TESTER'S"})
		}
	}

	return newController(
		"controllers.Test",
		views.TestView,
		c,
	).handler
}
