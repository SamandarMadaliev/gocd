package cmd

import (
	"log"

	"github.com/SamandarMadaliev/gocd/internal/commands"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [project_name] [module_path]",
	Args:  cobra.ExactArgs(2),
	Short: "Initialize a new gocd project",
	Long:  "Initialize a new gocd project and creates the project struct",
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		modulesPath := args[1]

		if projectName == "" {
			log.Fatal("Project name was not provided")
		}
		log.Println(projectName, "initializing")
		projectConfigs := commands.NewProjectConfigs(projectName, modulesPath)
		log.Println(
			"Project initialized successfully. You can start working on it now.",
		)
		log.Println(
			"To start working on the project, run the following command:",
		)
		commands.Init(projectConfigs)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
