package cmd

import (
	"github.com/mattvictorino/rating-api/api"
	"github.com/mattvictorino/rating-api/config"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Starts ratings API server",
	Long:  `Starts the ratings API server which provides endpoints to manage ratings.`,
	Run: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
		server := api.NewServer()
		server.Start()
	},
}
