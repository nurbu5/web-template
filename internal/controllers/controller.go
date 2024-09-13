package controllers

import (
	"context"
	"fmt"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"path/filepath"

	"github.com/nurbu5/web-template/internal/views"
)

type controllerFunc func(layouts *template.Template) http.HandlerFunc

type Controller struct {
	layouts    *template.Template
	name       string
	controller controllerFunc
}

func newController(layouts *template.Template, name string, files []string, controller controllerFunc, templateFS fs.FS) *Controller {
	layouts, err := layouts.Clone()

	c := &Controller{
		layouts:    layouts,
		name:       name,
		controller: controller,
	}

	if err != nil {
		c.throwInternalServerError("Failed to parse layout files", err)
	}

	fns := make([]string, len(files))
	for i, file := range files {
		fn := filepath.Join(views.TemplateDirectory, file)
		fns[i] = fn
	}

	if len(fns) > 0 {
		_, err = c.layouts.ParseFS(templateFS, fns...)

		if err != nil {
			c.throwInternalServerError("Failed to parse template files", err)
		}
	}

	return c
}

func (c *Controller) handler(w http.ResponseWriter, r *http.Request) {
	c.controller(c.layouts)(w, r)
}

func (c *Controller) throwInternalServerError(message string, err error) {
	ctx := context.Background()
	err = fmt.Errorf("controllers.%s: %w", c.name, err)
	slog.ErrorContext(ctx, message, slog.Any("err", err))
	slog.Error("controllers.Root", slog.Any("err", err))

	c.controller = func(_ *template.Template) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
