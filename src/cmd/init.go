package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "initialize the repository for VCS",
	Run:   runInit,
}

func init() {
	rootCmd.AddCommand(initCommand)
}

func runInit(cmd *cobra.Command, args []string) {
	fmt.Println("Running init command")
}
