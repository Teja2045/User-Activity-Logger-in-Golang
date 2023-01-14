package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra",
		Short: "iptracker cli app",
		Long:  `iptracker cli app`,
		Run: func(cmd *cobra.Command, args []string) {
			a := 5
			fmt.Println(a)
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
