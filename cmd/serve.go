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
	serveCmd.Flags().StringVar(&config.LocaleDefault, "locale", config.LocaleDefault, "default locale")
	serveCmd.Flags().StringSliceVar(&config.CollectionTypes, "collections", config.CollectionTypes, "collection type models")

	serveCmd.Flags().StringVar(&config.PostCmd, "cmd", config.PostCmd, "post commands to run")

	serveCmd.Flags().Int64Var(&config.DebouncedTimeout, "timeout", config.DebouncedTimeout, "debounced timeout")
	serveCmd.Flags().StringVar(&config.PostDebouncedCmd, "debounced-cmd", config.PostDebouncedCmd, "post debounced commands to run")

	// serveCmd.Flags().BoolVarP(&verbose, "verbose", "v", verbose, "post debounced commands to run")

	rootCmd.AddCommand(serveCmd)
}

// Start the server
func serve() error {
	helper.InitCommand(config.SiteDir, config.DebouncedTimeout)
	s := &service.Service{}
	return http.ListenAndServe(":8080", s)
}
