package structure

var mainFile = []Node{
	{
		Type:     File,
		Name:     "main.go",
		Template: "",
	},
}

var CMD = Node{
	Type:     Folder,
	Name:     "cmd",
	Template: "",
	Resource: mainFile,
}
