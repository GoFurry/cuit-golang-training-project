package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	addCmd := &cobra.Command{
		Use:   "add [numbers...]",
		Short: "Add integer numbers",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			total := 0
			for _, arg := range args {
				value, err := strconv.Atoi(arg)
				if err != nil {
					return err
				}
				total += value
			}
			fmt.Printf("sum = %d\n", total)
			return nil
		},
	}

	rootCmd.AddCommand(addCmd)
}
