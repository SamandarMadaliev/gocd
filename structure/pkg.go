package structure

import "github.com/SamandarMadaliev/gocd/internal/components"

var PKG = Node{
	Type:     Folder,
	Name:     "pkg",
	Template: "",
	Resource: pkgResources,
}

var pkgResources = []Node{
	databaseResources,
}

var databaseResources = Node{
	Type: Folder,
	Name: "database",
	Resource: []Node{
		postgres,
	},
}

var postgres = Node{
	Type: Folder,
	Name: "postgres",
	Resource: []Node{
		{
			Type:        File,
			Name:        "errors.go",
			Template:    "templates/pkg/database/postgres/errors.temp",
			ProcessFunc: components.MarshalPostgresPgxTemp,
		},
	},
}
