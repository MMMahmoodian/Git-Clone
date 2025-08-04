package cmd

import (
	"github.com/spf13/cobra"
)

var statusCommand = &cobra.Command{
	Use:   "status",
	Short: "Returns current status of repository",
	Run:   runStatus,
}

func init() {
	rootCmd.AddCommand(statusCommand)
}

func runStatus(cmd *cobra.Command, args []string) {
	cmd.Println("Running status command")
}
