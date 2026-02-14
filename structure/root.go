package structure

type NodeType string

const (
	File   NodeType = "file"
	Folder NodeType = "folder"
)

type Node struct {
	Type     NodeType
	Name     string
	Template string
	Resource []Node
}

var ignoreFile = Node{
	Type:     File,
	Name:     ".gitignore",
	Template: "templates/root/gitignore.temp",
	Resource: make([]Node, 0),
}

var ProjectStruct = Node{
	Type: Folder,
	Resource: []Node{
		CMD,
		INTERNAL,
		PKG,
		ignoreFile,
	},
}
