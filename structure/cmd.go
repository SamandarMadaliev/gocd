package structure

var mainFile = []Node{
	{
		Type:     File,
		Name:     "main.go",
		Template: "templates/cmd/main.temp",
	},
}

var CMD = Node{
	Type:     Folder,
	Name:     "cmd",
	Template: "",
	Resource: mainFile,
}
