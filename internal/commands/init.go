package commands

import (
	"io"
	"os"
)

func Init() {
	rootPath := "./project/"
	var projectFolders = []string{
		rootPath + "cmd",
		rootPath + "internal",
		rootPath + "pkg",
	}

	var projectFiles = map[string]string{
		"./internal/templates/cmd/cmd_init.temp": rootPath + "cmd/init.go",
		"./internal/templates/cmd/cmd_main.temp": rootPath + "cmd/main.go",
	}

	// Generate folders
	for _, dir := range projectFolders {
		generateFolder(dir, 0755)
	}

	for fromFilePath, toFilePath := range projectFiles {
		generateFileFromTemplate(fromFilePath, toFilePath)
	}
}

func generateFolder(fullPath string, permission os.FileMode) error {
	return os.MkdirAll(fullPath, permission)
}

func generateFileFromTemplate(fromFile string, toFile string) error {
	// Open source file
	src, err := os.Open(fromFile)
	if err != nil {
		return err
	}
	defer src.Close()

	// Create destination file
	dst, err := os.Create(toFile)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy content from source to destination
	_, err = io.Copy(dst, src)
	return err
}
