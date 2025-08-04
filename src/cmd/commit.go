package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var commitCommand = &cobra.Command{
	Use:   "commit",
	Short: "Commit staged files to VCS",
	Run:   runCommit,
}

func init() {
	rootCmd.AddCommand(commitCommand)
}

func runCommit(cmd *cobra.Command, args []string) {
	fmt.Println("Running commit command")
}
