package cmd

import (
	"github.com/spf13/cobra"

	"strapiwebhook/helper"
	"strapiwebhook/service/config"
)

// Serve gRPC and optional servers: gRPC web proxy + HTTP
var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Write sample templates",
	RunE: func(cmd *cobra.Command, args []string) error {
		return helper.WriteAllEmbed(config.TemplateDir)
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)
}
