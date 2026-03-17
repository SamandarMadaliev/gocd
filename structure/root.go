package structure

import "text/template"

type NodeType string

const (
	File   NodeType = "file"
	Folder NodeType = "folder"
)

type Node struct {
	Type        NodeType
	Name        string
	Template    string
	ProcessFunc func(source *template.Template, values any) ([]byte, error)
	Resource    []Node
}

var ignoreFile = Node{
	Type:     File,
	Name:     ".gitignore",
	Template: "templates/root/gitignore.temp",
}

var dockerCompose = Node{
	Type:     File,
	Name:     "docker-compose.yml",
	Template: "templates/root/docker-compose.temp",
}

var envFile = Node{
	Type:     File,
	Name:     ".env",
	Template: "templates/root/env.temp",
}

var airTomp = Node{
	Type:     File,
	Name:     ".air.toml",
	Template: "templates/root/air.temp",
}

var dockerfile = Node{
	Type:     File,
	Name:     "Dockerfile",
	Template: "templates/root/dockerfile.temp",
}

var dockerfileDev = Node{
	Type:     File,
	Name:     "Dev.Dockerfile",
	Template: "templates/root/dockerfile-dev.temp",
}

var ProjectStruct = Node{
	Type: Folder,
	Resource: []Node{
		CMD,
		INTERNAL,
		MIGRATIONS,
		PKG,
		ignoreFile,
		dockerCompose,
		envFile,
		airTomp,
		dockerfile,
		dockerfileDev,
	},
}
