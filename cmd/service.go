package cmd

import (
	"appgraph/cmd/service"

	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:     "service",
	Aliases: []string{"svc", "s"},
	Short:   "Register, update, delete a service in AppGraph",
	Long:    ``,
}

func init() {
	serviceCmd.AddCommand(service.CreateCmd)
	serviceCmd.AddCommand(service.ListCmd)
	serviceCmd.AddCommand(service.GenerateSbomCmd)
}
