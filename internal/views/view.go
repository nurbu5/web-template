package views

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

//go:embed web/templates
var templateFS embed.FS
var layouts = template.Must(template.ParseFS(templateFS, "web/templates/layouts/*.html"))

type ViewProducer func() (*View, error)

const TemplateDirectory = "web/templates"

type View struct {
	name     string
	layout   string
	template *template.Template
}

func newView(name, layout string, files []string) (*View, error) {
	l, err := layouts.Clone()

	if err != nil {
		return nil, fmt.Errorf("views.newView (name: %s) (Failed to parse layout files): %w", name, err)
	}

	fns := make([]string, len(files))
	for i, file := range files {
		fn := filepath.Join(TemplateDirectory, file)
		fns[i] = fn
	}

	if len(fns) > 0 {
		l, err = l.ParseFS(templateFS, fns...)

		if err != nil {
			return nil, fmt.Errorf("views.newView (name: %s) (Failed to parse template files): %w", name, err)
		}
	}

	return &View{
		name:     name,
		layout:   layout,
		template: l,
	}, nil
}

func (v *View) Execute(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return v.template.Lookup(v.layout).Execute(w, data)
}
