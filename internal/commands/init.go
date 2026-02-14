package commands

import (
	"log"
	"os"

	"github.com/SamandarMadaliev/gocd/assets"
	"github.com/SamandarMadaliev/gocd/structure"
)

func Init(projectName string) {
	generateProject(structure.ProjectStruct, projectName)
}

func generateProject(node structure.Node, rootPath string) {
	if node.Type == structure.Folder {
		err := generateFolder(rootPath+node.Name, 0755)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if node.Type == structure.File {
		err := generateFileFromTemplate(node.Template, rootPath+node.Name)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if len(node.Resource) != 0 {
		for _, nodeResource := range node.Resource {
			generateProject(nodeResource, rootPath+node.Name+"/")
		}
	}
}

func generateFolder(fullPath string, permission os.FileMode) error {
	return os.MkdirAll(fullPath, permission)
}

func generateFileFromTemplate(fromFile string, toFile string) error {
	src, err := assets.GetTemplate(fromFile)
	if err != nil {
		return err
	}

	// Create destination file
	dst, err := os.Create(toFile)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy content from source to destination
	return os.WriteFile(toFile, src, 0644)
}
