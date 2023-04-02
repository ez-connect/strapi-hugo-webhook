package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"strapiwebhook/service/config"
	"strapiwebhook/service/rest"
)

var (
	fetchModelName    string
	fetchEntryEnpoint string
	fetchEntryId      string
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch entries from Strapi Rest API",
	Run: func(cmd *cobra.Command, args []string) {
		panic("Not impl")
	},
}

var fetchEntryListCmd = &cobra.Command{
	Use:   "list",
	Short: "fetch entries from Strapi Rest API",
	Run: func(cmd *cobra.Command, args []string) {
		uri := fmt.Sprintf("%s/%s", config.StrapiAddr, fetchEntryEnpoint)
		err := rest.FetchAndWriteEntryList(config.SiteDir, config.TemplateDir, fetchModelName, uri, config.StrapiToken)
		if err != nil {
			panic(err)
		}
	},
}

var fetchEntryCmd = &cobra.Command{
	Use:   "get",
	Short: "Fetch an entry from Strapi Rest API",
	Run: func(cmd *cobra.Command, args []string) {
		uri := fmt.Sprintf("%s/%s/%s", config.StrapiAddr, fetchEntryEnpoint, fetchEntryId)
		err := rest.FetchAndWriteEntry(config.SiteDir, config.TemplateDir, fetchModelName, uri, config.StrapiToken)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	fetchCmd.Flags().StringVarP(&fetchModelName, "model", "m", fetchModelName, "entry model")
	fetchCmd.Flags().StringVarP(&fetchEntryEnpoint, "endpoint", "e", fetchEntryEnpoint, "entry enpoint")
	fetchEntryCmd.Flags().StringVar(&fetchEntryId, "id", fetchEntryId, "entry ID")

	fetchCmd.AddCommand(fetchEntryListCmd)
	fetchCmd.AddCommand(fetchEntryCmd)

	rootCmd.AddCommand(fetchCmd)
}
