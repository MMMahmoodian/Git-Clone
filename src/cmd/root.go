package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "cli handler",
	Short: "Simple VCS cli tool",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
