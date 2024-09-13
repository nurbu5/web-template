package views

import (
	"html/template"
	"io/fs"
)

const TemplateDirectory = "./web/templates"

// ConfigViews parses the layout files.
func ConfigViews(templateFS fs.FS) *template.Template {
	return template.Must(
		template.ParseFS(templateFS, "web/templates/layouts/*.html"),
	)
}
