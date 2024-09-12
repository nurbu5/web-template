package views

import (
	"html/template"
	"path/filepath"
)

const LayoutDirectory = "./web/templates/layouts"
const TemplateDirectory = "./web/templates"

// ConfigViews parses the layout files.
func ConfigViews() *template.Template {
	files, err := filepath.Glob(LayoutDirectory + "/*")
	if err != nil {
		panic(err)
	}

	return template.Must(template.ParseFiles(files...))
}
