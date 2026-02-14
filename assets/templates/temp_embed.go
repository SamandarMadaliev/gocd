package templates

import (
	"embed"
	"fmt"
)

//go:embed **/*.temp
var fs embed.FS

func GetTemplate(file string) ([]byte, error) {
	fmt.Println(fs.ReadDir("cmd/"))
	readFile, err := fs.ReadFile("./" + file)
	if err != nil {
		return nil, err
	}
	return readFile, nil
}
