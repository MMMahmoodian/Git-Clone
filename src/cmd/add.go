package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add files to staging area of VCS",
	Run:   runAdd,
}

func init() {
	rootCmd.AddCommand(addCommand)
}

func runAdd(cmd *cobra.Command, args []string) {
	fmt.Println("Running add command")
}
