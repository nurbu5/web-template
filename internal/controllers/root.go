package controllers

import (
	"net/http"

	"github.com/nurbu5/web-template/internal/views"
)

// Path: /
func Root() http.HandlerFunc {
	c := func(v *views.View) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			v.Execute(w, nil)
		}
	}

	return newController(
		"controllers.Root",
		views.RootView,
		c,
	).handler()
}
