package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rating-api",
	Short: "Rating API is a service for managing ratings",
	Long: `Rating API is a service for managing ratings. 
	It provides endpoints to create, read, update, and delete ratings.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
