package service

import (
	"appgraph/db"
	"context"

	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Request the creation of a service identity in AppGraph",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		dbClient, err := db.NewPostgres(ctx)
		if err != nil {
			return err
		}

		_, err = dbClient.CreateService(ctx, &db.Service{
			Name:       args[0],
			Type:       serviceType,
			Csi:        csi,
			Status:     "Draft",
			Repository: repo,
		})
		if err != nil {
			return err
		}

		return nil
	},
}

var (
	csi         string
	labels      []string
	serviceType string
	repo        string
)

func init() {
	CreateCmd.Flags().StringVar(&csi, "csi", "", "Enter your CSI")
	CreateCmd.Flags().StringArrayVar(&labels, "add-label", []string{}, "Add labels for your service")
	CreateCmd.Flags().StringVar(&serviceType, "type", "component", "Choose a service type ; 'component' or 'infra'")
	CreateCmd.Flags().StringVar(&repo, "repo", "", "Enter the repository link to your service")
	CreateCmd.MarkFlagRequired("csi")
	CreateCmd.MarkFlagRequired("repo")
}
