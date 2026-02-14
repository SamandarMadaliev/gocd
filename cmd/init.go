package cmd

import (
	"log"

	"github.com/SamandarMadaliev/gocd/internal/commands"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [project_name]",
	Args:  cobra.ExactArgs(1),
	Short: "Initialize a new gocd project",
	Long:  "Initialize a new gocd project and creates the project struct",
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		if projectName == "" {
			log.Fatal("Project name was not provided")
		}
		log.Println(projectName, "initializing")

		commands.Init(projectName)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
