package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	var name string

	greetCmd := &cobra.Command{
		Use:   "greet",
		Short: "Print a greeting message",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" {
				name = "Go learner"
			}
			fmt.Printf("hello, %s\n", name)
		},
	}

	greetCmd.Flags().StringVarP(&name, "name", "n", "", "name to greet")
	rootCmd.AddCommand(greetCmd)
}
