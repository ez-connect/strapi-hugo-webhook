package cmd

import (
	"github.com/spf13/cobra"

	"strapiwebhook/helper"
)

var (
	tplOutDir = "tpl"
)

// Serve gRPC and optional servers: gRPC web proxy + HTTP
var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Write sample templates",
	RunE: func(cmd *cobra.Command, args []string) error {
		return helper.WriteAllEmbed(tplOutDir)
	},
}

func init() {
	templateCmd.Flags().StringVar(&tplOutDir, "dir", tplOutDir, "the output dir")
	rootCmd.AddCommand(templateCmd)
}
