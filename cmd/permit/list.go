package permit

import (
	"appgraph/db"
	"appgraph/output"
	"context"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:     "list",
	Short:   "Get registered permits from AppGraph for a CSI",
	Aliases: []string{"ls"},
	Long:    ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		dbClient, err := db.NewPostgres(ctx)
		if err != nil {
			return err
		}

		permits, err := dbClient.ListPermits(ctx, csi)
		if err != nil {
			return err
		}

		output.PrintPermitsTable(permits)

		return nil
	},
}

func init() {
	ListCmd.Flags().StringVar(&csi, "csi", "", "Enter your CSI")
	ListCmd.MarkFlagRequired("csi")
}
