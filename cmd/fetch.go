package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"strapiwebhook/service/config"
	"strapiwebhook/service/rest"
)

var (
	fetchModel   = "section"
	fetchEnpoint = "sections"
	fetchPage    = 1
	fetchSize    = 100
	fetchId      = 1
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch entries from Strapi Rest API",
}

var fetchEntryListCmd = &cobra.Command{
	Use:   "list",
	Short: "fetch entries from Strapi Rest API",
	RunE: func(cmd *cobra.Command, args []string) error {
		uri := fmt.Sprintf(
			"%s/api/%s?populate=*&pagination[page]=%v&pagination[pageSize]=%v",
			config.StrapiAddr,
			fetchEnpoint,
			fetchPage,
			fetchSize,
		)

		fmt.Println("fetch:", uri)
		return rest.FetchAndWriteEntryList(config.SiteDir, config.TemplateDir, fetchModel, uri, config.StrapiToken)
	},
}

var fetchEntryCmd = &cobra.Command{
	Use:   "get",
	Short: "Fetch an entry from Strapi Rest API",
	RunE: func(cmd *cobra.Command, args []string) error {
		uri := fmt.Sprintf("%s/api/%s/%v", config.StrapiAddr, fetchEnpoint, fetchId)
		fmt.Println("fetch:", uri)
		return rest.FetchAndWriteEntry(config.SiteDir, config.TemplateDir, fetchModel, uri, config.StrapiToken)
	},
}

func init() {
	fetchCmd.PersistentFlags().StringVarP(&fetchModel, "model", "m", fetchModel, "entry model")
	fetchCmd.PersistentFlags().StringVarP(&fetchEnpoint, "endpoint", "e", fetchEnpoint, "entry enpoint")

	fetchEntryListCmd.Flags().IntVarP(&fetchPage, "page", "p", fetchPage, "current page")
	fetchEntryListCmd.Flags().IntVarP(&fetchSize, "size", "S", fetchSize, "page size")
	fetchEntryCmd.Flags().IntVar(&fetchId, "id", fetchId, "entry ID")

	fetchCmd.AddCommand(fetchEntryListCmd)
	fetchCmd.AddCommand(fetchEntryCmd)

	rootCmd.AddCommand(fetchCmd)
}
