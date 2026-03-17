package commands

import (
	"log"
	"os"

	"github.com/SamandarMadaliev/gocd/assets"
	"github.com/SamandarMadaliev/gocd/internal/components"
	"github.com/SamandarMadaliev/gocd/pkg/helpers"
	"github.com/SamandarMadaliev/gocd/structure"
)

type ProjectConfigs struct {
	ProjectName string
	ModulePath  string
}

func NewProjectConfigs(projectName string, modulePath string) *ProjectConfigs {
	return &ProjectConfigs{
		ProjectName: projectName,
		ModulePath:  modulePath,
	}
}

func Init(projectConfigs *ProjectConfigs) {
	if len([]rune(projectConfigs.ProjectName)) > 1 && projectConfigs.ProjectName != "." {
		exists, err := helpers.FolderExists(projectConfigs.ProjectName)
		if err != nil {
			log.Fatalln("Error checking project existence:", err.Error())
		}
		if exists {
			log.Fatalln("Project already exists")
		}
	}

	generateProject(structure.ProjectStruct, projectConfigs.ProjectName)
}

func generateProject(node structure.Node, rootPath string) {
	if node.Type == structure.Folder {
		err := generateFolder(rootPath+node.Name, 0755)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if node.Type == structure.File {
		err := generateFileFromTemplate(node, rootPath)
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

func generateFileFromTemplate(node structure.Node, rootPath string) error {
	temp, err := assets.GetTemplate(node.Template)
	if err != nil {
		return err
	}

	data := components.PostgresPgxPlaceholders{
		ProjectPlaceholders: components.ProjectPlaceholders{
			ProjectName: node.Name,
		},
	}
	src, err := node.ProcessFunc(temp, data)

	// Create destination file
	dst, err := os.Create(rootPath + node.Name)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy content from source to destination
	return os.WriteFile(rootPath+node.Name, src, 0644)
}
