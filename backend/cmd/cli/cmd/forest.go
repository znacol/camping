package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// forestCmd represents the forest command
var forestCmd = &cobra.Command{
	Use:   "forest",
	Short: "Manage forests",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("forest called")
	},
}

func init() {
	rootCmd.AddCommand(forestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// forestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// forestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
