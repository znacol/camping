package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/znacol/camping/backend/cmd/cli/api"
	pb "github.com/znacol/camping/backend/proto"
)

// siteCmd represents the site command
var siteCmd = &cobra.Command{
	Use:   "site",
	Short: "Manage camping sites",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("site called")
	},
}

func init() {
	rootCmd.AddCommand(siteCmd)

	siteCmd.AddCommand(createSite())
}

func createSite() *cobra.Command {
	site := &pb.Site{}
	req := &pb.CreateSiteRequest{
		Site: site,
	}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create new site",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := api.CampingManager.CreateSite(context.Background(), req)

			return err
		},
	}
	cmd.Flags().Float32VarP(&site.Latitude, "latitude", "c", 0, "site's latitude")
	cmd.Flags().Float32VarP(&site.Longitude, "longitude", "l", 0, "site's longitude")
	cmd.Flags().Int64VarP(&site.NationalForestId, "nationalforest", "f", 1, "national forest ID")
	cmd.Flags().Int64VarP(&site.DistrictId, "district", "d", 1, "district ID")
	cmd.Flags().Int64VarP(&site.Altitude, "altitude", "a", 0 , "site altitude")
	return cmd
}

