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
