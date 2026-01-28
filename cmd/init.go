package cmd

import (
	"github.com/SamandarMadaliev/gocd/internal/commands"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new gocd project",
	Long:  "Initialize a new gocd project and creates the project struct",
	Run: func(cmd *cobra.Command, args []string) {
		commands.Init()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
