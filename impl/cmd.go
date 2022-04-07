package impl

import (
	"flag"
	"fmt"
	"os"

	"github.com/oklog/run"

	"strapi-webhook/base"
)

var (
	siteDir       string // hugo site dir
	strapiAddr    string // strapi server
	localeDefault string // default locale
	gitCommitMsg  string // git commit message, leave blank to ignore git commit & push
	gitTimeout    int64  // git timeout in seconds
)

// Adds flags & print flags help used `flag.FlagSet.Usage`
func UpdateFlagSet(fs *flag.FlagSet) {
	fs.StringVar(&strapiAddr, "s", "http://localhost:1337", "strapi listen address")
	fs.StringVar(&localeDefault, "l", "en", "default locale")
	fs.StringVar(&gitCommitMsg, "m", "", "git commit message, leave blank to ignore")
	fs.Int64Var(&gitTimeout, "t", 300, "git timeout in second")

	fs.Usage = func() {
		fmt.Println(base.Name, fmt.Sprintf("v%s - %s", base.Version, base.Description))
		fmt.Println("USAGE:", base.Name, "[OPTIONS]")
		fmt.Println("\nOPTIONS")
		fs.PrintDefaults()
	}
}

// Parses args from 'fs`` or add a an actor to the group `g`
func AddToCmd(fs *flag.FlagSet, g *run.Group) {
	// No args, print usage then exit
	if fs.NArg() < 1 {
		fs.Usage()
		os.Exit(1)
	}

	siteDir = fs.Arg(0)

	// Set Strapi + Hugo site dir + git message
	SetStrapiAddr(strapiAddr)
	SetSiteDir(siteDir)
	SetDefaultLocale(localeDefault)
	SetGit(gitCommitMsg, gitTimeout)
}
