package cmd

import (
	"net/http"

	"github.com/spf13/cobra"

	"strapiwebhook/helper"
	"strapiwebhook/service"
	"strapiwebhook/service/config"
)

// var verbose = false

// Serve gRPC and optional servers: gRPC web proxy + HTTP
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return serve()
	},
}

// Return the serve command
func GetServeCmd() *cobra.Command {
	return serveCmd
}

func init() {
	serveCmd.Flags().StringVarP(&config.Cmd, "cmd", "c", config.Cmd, "commands to run after trigger")

	// Debounced commands
	serveCmd.Flags().StringVar(&config.DebouncedCmd, "debounced-cmd", config.DebouncedCmd, "post debounced commands to run")
	serveCmd.Flags().Int64Var(&config.DebouncedTimeout, "debounced-timeout", config.DebouncedTimeout, "debounced timeout in second")

	rootCmd.AddCommand(serveCmd)
}

// Start the server
func serve() error {
	helper.InitCommand(config.SiteDir, config.DebouncedTimeout)
	s := &service.Service{}
	return http.ListenAndServe(":8080", s)
}
