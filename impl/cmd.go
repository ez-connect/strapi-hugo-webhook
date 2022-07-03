package impl

import (
	"github.com/spf13/cobra"
)

const (
	subscriptionAPI = "api/subscriptions?filters[topics][$contains]=<placeholder>&fields=email&pagination[pageSize]=100"
)

var (
	strapiAddr  = "http://localhost:1337" // strapi server
	strapiToken = ""                      // strapi token

	siteDir         = "web"                                    // hugo site dir
	localeDefault   = "en"                                     // default locale
	singleTypes     = []string{"site", "home", "nav", "about"} // single type models, comma separated
	collectionTypes = []string{"contributor", "article", "document", "career", "project", "page", "resume"}

	gitCommitMsg = ""         // git commit message, leave blank to ignore git commit & push
	gitTimeout   = int64(300) // git timeout in seconds
)

// Add extra `flags` to `serve` commands
func AddCmd(serveCmd *cobra.Command) {
	// Add extra flags
	serveCmd.Flags().StringVar(&strapiAddr, "strapi", strapiAddr, "strapi listen address")
	serveCmd.Flags().StringVar(&strapiToken, "token", strapiToken, "strapi api token")

	serveCmd.Flags().StringVar(&siteDir, "dir", siteDir, "hugo site dir")
	serveCmd.Flags().StringVar(&localeDefault, "locale", localeDefault, "default locale")
	serveCmd.Flags().StringSliceVar(&singleTypes, "single", singleTypes, "single type models")
	serveCmd.Flags().StringSliceVar(&collectionTypes, "collection", collectionTypes, "single type models")

	serveCmd.Flags().StringVar(&gitCommitMsg, "commit", gitCommitMsg, "git commit message, leave blank to ignore")
	serveCmd.Flags().Int64Var(&gitTimeout, "timeout", gitTimeout, "git timeout in second")

	// Override command
	// fn := serveCmd.Run
	// serveCmd.Run = func(cmd *cobra.Command, args []string) {
	// 	// Do before the server start...
	// 	fn(cmd, args)
	// }
}
