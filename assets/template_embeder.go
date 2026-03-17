package assets

import (
	"embed"
	"text/template"
)

//go:embed templates
var templates embed.FS

func GetTemplate(file string) (*template.Template, error) {
	readFile, err := template.ParseFS(templates, file)
	if err != nil {
		return nil, err
	}
	return readFile, nil
}
