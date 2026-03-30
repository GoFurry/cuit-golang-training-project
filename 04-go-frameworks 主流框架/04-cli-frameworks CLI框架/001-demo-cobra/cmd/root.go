package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "framework-cli",
	Short: "A small Cobra demo for Go learners",
	Long:  "framework-cli demonstrates subcommands, flags, and argument parsing with Cobra.",
}

func Execute() error {
	return rootCmd.Execute()
}
