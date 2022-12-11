package cmd

import (
	"net/http"

	"github.com/spf13/cobra"

	"strapiwebhook/helper"
	"strapiwebhook/service"
	"strapiwebhook/service/config"
)

// Serve gRPC and optional servers: gRPC web proxy + HTTP
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

// Return the serve command
func GetServeCmd() *cobra.Command {
	return serveCmd
}

func init() {
	// Add extra flags
	serveCmd.Flags().StringVar(&config.StrapiAddr, "strapi", config.StrapiAddr, "strapi listen address")
	serveCmd.Flags().StringVar(&config.StrapiToken, "token", config.StrapiToken, "strapi api token")

	serveCmd.Flags().StringVar(&config.SiteDir, "site-dir", config.SiteDir, "website site root dir")
	serveCmd.Flags().StringVar(&config.LocaleDefault, "locale", config.LocaleDefault, "default locale")
	serveCmd.Flags().StringSliceVar(&config.CollectionTypes, "collections", config.CollectionTypes, "collection type models")

	serveCmd.Flags().StringVar(&config.TemplateDir, "template", config.TemplateDir, "template dir")
	serveCmd.Flags().StringVar(&config.PostCmd, "cmd", config.PostCmd, "post commands to run")

	serveCmd.Flags().Int64Var(&config.DebouncedTimeout, "timeout", config.DebouncedTimeout, "debounced timeout")
	serveCmd.Flags().StringVar(&config.PostDebouncedCmd, "debounced-cmd", config.PostDebouncedCmd, "post debounced commands to run")

	rootCmd.AddCommand(serveCmd)
}

// Start the server
func serve() {
	helper.InitCommand(config.SiteDir, config.DebouncedTimeout)
	s := &service.Service{}
	http.ListenAndServe(":8080", s)
}
