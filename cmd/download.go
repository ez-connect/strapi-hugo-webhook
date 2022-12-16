package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"strapiwebhook/helper"
)

// Serve gRPC and optional servers: gRPC web proxy + HTTP
var downloadCmd = &cobra.Command{
	Use:   "template",
	Short: "Write sample templates",
	Run: func(cmd *cobra.Command, args []string) {
		err := helper.WriteAllEmbed(tplOutDir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}
