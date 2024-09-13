package controllers

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/nurbu5/web-template/internal/views"
)

type controllerFunc func(v *views.View) http.HandlerFunc

type Controller struct {
	name    string
	view    *views.View
	control controllerFunc
}

func newController(n string, vp views.ViewProducer, c controllerFunc) *Controller {
	v, err := vp()
	out := &Controller{
		name:    n,
		view:    v,
		control: c,
	}

	if err != nil {
		out.throwInternalServerError("controllers.newController", err)
	}

	return out
}

func (c *Controller) handler(w http.ResponseWriter, r *http.Request) {
	c.control(c.view)(w, r)
}

func (c *Controller) throwInternalServerError(message string, err error) {
	ctx := context.Background()
	err = fmt.Errorf("controllers.%s: %w", c.name, err)
	slog.ErrorContext(ctx, message, slog.Any("err", err))
	slog.Error("controllers.Root", slog.Any("err", err))

	c.control = func(_ *views.View) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
