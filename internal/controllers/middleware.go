package controllers

import (
	"net/http"

	"github.com/NYTimes/gziphandler"
)

type middleware func(http.Handler) http.Handler

func middlewareList() []middleware {
	// You can add middleware here (called in order)
	return []middleware{
		gziphandler.GzipHandler,
	}
}
