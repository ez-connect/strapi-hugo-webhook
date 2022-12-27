package cmd

import (
	"github.com/spf13/cobra"
)

// Serve gRPC and optional servers: gRPC web proxy + HTTP
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch entries from Strapi Rest API",
	Run: func(cmd *cobra.Command, args []string) {
		panic("Not impl")
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
