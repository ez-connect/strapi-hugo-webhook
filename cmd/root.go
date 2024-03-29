package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"

	"strapiwebhook/helper/zlog"
	"strapiwebhook/service/config"
)

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "strapiwebhook",
	Short:   "Strapi webhook service",
	Version: fmt.Sprintf("v%s-%s/%s %s/%s BuildDate=%s\n", config.Version, config.Hash, config.Branch, runtime.GOOS, runtime.GOARCH, config.BuildDate),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Return the root command
func GetRootCmd() *cobra.Command {
	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gkgen.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVarP(&config.SiteDir, "site-dir", "d", config.SiteDir, "website site root dir")
	rootCmd.PersistentFlags().StringVarP(&config.TemplateDir, "template-dir", "T", config.TemplateDir, "template dir")

	rootCmd.PersistentFlags().StringVarP(&config.LocaleDefault, "locale", "l", config.LocaleDefault, "default locale")
	rootCmd.PersistentFlags().StringSliceVarP(&config.CollectionTypes, "collections", "C", config.CollectionTypes, "collection type models")
	rootCmd.PersistentFlags().StringSliceVarP(&config.SingleTypes, "singles", "S", config.SingleTypes, "single type models")

	rootCmd.PersistentFlags().StringVarP(&config.StrapiAddr, "strapi-host", "s", config.StrapiAddr, "strapi listen address")
	rootCmd.PersistentFlags().StringVarP(&config.StrapiToken, "strapi-token", "t", config.StrapiToken, "strapi api token")

	// Init the logger before used
	zlog.InitLogger(config.BuildMode == "production")
}
