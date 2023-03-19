package cmd

import (
	"appgraph/cmd/permit"

	"github.com/spf13/cobra"
)

var permitsCmd = &cobra.Command{
	Use:     "permit",
	Aliases: []string{"prt", "p"},
	Short:   "Request a build or deploy permit",
	Long:    ``,
}

func init() {
	permitsCmd.AddCommand(permit.RegisterCmd)
	permitsCmd.AddCommand(permit.ListCmd)
}
