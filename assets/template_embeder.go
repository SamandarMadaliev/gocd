package assets

import (
	"embed"
)

//go:embed templates
var templates embed.FS

func GetTemplate(file string) ([]byte, error) {
	readFile, err := templates.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return readFile, nil
}
