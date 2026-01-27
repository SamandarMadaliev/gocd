package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new gocd project",
	Long:  "Initialize a new gocd project and creates the project struct",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initializing gocd project...")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
