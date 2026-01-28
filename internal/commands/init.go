package commands

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/SamandarMadaliev/gocd/internal/types"
)

func Init() {
	projectStructure := types.NestedMap{
		"project": types.NestedMap{
			"cmd": []types.NestedMap{
				{
					"file_name":     "main.go",
					"template_path": "./internal/templates/cmd/cmd_main.temp",
				},
				{
					"file_name":     "init.go",
					"template_path": "./internal/templates/cmd/cmd_init.temp",
				},
			},
		},
	}
	fmt.Println(projectStructure["project"])
	//for key, value := range projectStructure {
	//	fmt.Println(key, value)
	//}
	generateFolderStructure(projectStructure)
	fmt.Println("Initializing gocd project...")
}

func generateFolderStructure(projectStructureMap types.NestedMap) {
	for key, value := range projectStructureMap {
		if key != "" {
			err := createFolder(key, "./", 0755)
			if err != nil {
				log.Fatal(err)
			}
		}

		if innerFolder, ok := value.(types.NestedMap); ok {
			generateFolderStructure(innerFolder)
		}

		if innerFiles, ok := value.([]types.NestedMap); ok {
			for _, file := range innerFiles {
				err := createFileFromTemplate(file["template_path"].(string), "./"+file["file_name"].(string), 0644)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

	}
}

func createFolder(folderName string, path string, permission os.FileMode) error {
	return os.Mkdir(path+folderName, permission)
}

func createFileFromTemplate(fromFile string, toFile string, toFilePermission os.FileMode) error {
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
