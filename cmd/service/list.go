package service

import (
	"appgraph/db"
	"appgraph/output"
	"context"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:     "list",
	Short:   "Get a list of registered services from AppGraph",
	Long:    ``,
	Aliases: []string{"ls"},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		dbClient, err := db.NewPostgres(ctx)
		if err != nil {
			return err
		}

		listOfServices, err := dbClient.ListService(ctx, csi)
		if err != nil {
			return err
		}

		output.PrintServiceTable(listOfServices)

		return nil
	},
}

func init() {
	ListCmd.Flags().StringVar(&csi, "csi", "", "Enter your CSI")
	ListCmd.MarkFlagRequired("csi")
}
