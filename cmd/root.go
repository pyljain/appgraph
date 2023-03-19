package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ag",
	Short: "ag is a CLI to communicate with the AppGraph",
	Long:  ``,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(serviceCmd)
	rootCmd.AddCommand(permitsCmd)
}
