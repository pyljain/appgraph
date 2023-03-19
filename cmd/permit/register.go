package permit

import (
	"appgraph/db"
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var permitType string
var csi string

var RegisterCmd = &cobra.Command{
	Use:     "register",
	Short:   "Register a permit with AppGraph",
	Aliases: []string{"rg"},
	Long:    ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		dbClient, err := db.NewPostgres(ctx)
		if err != nil {
			return err
		}

		var permitId int
		if permitType == "build" {
			permitId, err = createPermitToBuild(args, dbClient, ctx)
		} else if permitType == "deploy" {
			permitId, err = createPermitToDeploy(args, dbClient, ctx)
		}
		if err != nil {
			return err
		}

		log.Printf("Permit created %d", permitId)

		return nil
	},
}

func createPermitToBuild(args []string, dbClient *db.Postgres, ctx context.Context) (int, error) {
	serviceArgsInt := []int{}

	for _, s := range args {
		i, err := strconv.Atoi(s)
		if err != nil {
			return -1, err
		}

		serviceArgsInt = append(serviceArgsInt, i)
	}

	permit, err := dbClient.RegisterPermit(ctx, &db.Permit{
		Type:     permitType,
		Csi:      csi,
		Status:   "Requested",
		Services: serviceArgsInt,
	})
	if err != nil {
		return -1, err
	}

	return permit, nil
}

func createPermitToDeploy(args []string, dbClient *db.Postgres, ctx context.Context) (int, error) {

	// Get a list of services for a permit
	permitId, err := strconv.Atoi(args[0])
	if err != nil {
		return -1, err
	}

	services, err := dbClient.GetServicesForPermit(ctx, permitId)
	if err != nil {
		return -1, err
	}
	log.Printf("createPermitToDeploy Services: %v", services)

	// Get ids of services
	serviceIds := []int{}
	for _, s := range services {
		serviceIds = append(serviceIds, s.Id)
	}

	permit, err := dbClient.RegisterPermit(ctx, &db.Permit{
		Type:     permitType,
		Csi:      csi,
		Status:   "Requested",
		Services: serviceIds,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Clone repositories for each service
	for _, s := range services {
		// Get path to repository
		components := strings.Split(s.Repository, "/")
		repoName := components[len(components)-1]
		p := path.Join(os.TempDir(), repoName)
		log.Printf("Path: %s", p)

		_, err := git.PlainClone(p, false, &git.CloneOptions{
			URL: s.Repository,
		})
		if err != nil {
			return -1, err
		}

		fmt.Printf("Running validator on %s\n", s.Repository)
	}

	return permit, nil
}

func init() {
	RegisterCmd.Flags().StringVar(&csi, "csi", "", "Enter your CSI")
	RegisterCmd.Flags().StringVar(&permitType, "type", "build", "Choose a permit type ; 'build' or 'deploy'")
	RegisterCmd.MarkFlagRequired("csi")
}
