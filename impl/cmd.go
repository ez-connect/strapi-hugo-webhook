package impl

import (
	"github.com/spf13/cobra"
)

var (
	siteDir       string // hugo site dir
	strapiAddr    string // strapi server
	localeDefault string // default locale
	gitCommitMsg  string // git commit message, leave blank to ignore git commit & push
	gitTimeout    int64  // git timeout in seconds
)

// Add extra `flags` to `serve` commands
func AddCmd(serveCmd *cobra.Command) {
	// Add extra flags
	serveCmd.Flags().StringVar(&strapiAddr, "s", "http://localhost:1337", "strapi listen address")
	serveCmd.Flags().StringVar(&localeDefault, "l", "en", "default locale")
	serveCmd.Flags().StringVar(&gitCommitMsg, "m", "", "git commit message, leave blank to ignore")
	serveCmd.Flags().Int64Var(&gitTimeout, "t", 300, "git timeout in second")

	// Override command
	fn := serveCmd.Run
	serveCmd.Run = func(cmd *cobra.Command, args []string) {
		// nolint:errcheck
		GetLogger().Log("strapi", strapiAddr, "locale", localeDefault, "commit", gitCommitMsg, "timeout", gitTimeout)
		SetStrapiAddr(strapiAddr)
		SetSiteDir(siteDir)
		SetDefaultLocale(localeDefault)
		SetGit(gitCommitMsg, gitTimeout)

		fn(cmd, args)
	}
}
