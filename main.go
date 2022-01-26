package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-kit/log"
	"github.com/oklog/run"

	"strapi-webhook/base"
	"strapi-webhook/base/server"
	"strapi-webhook/impl"
)

func main() {
	// Define our flags. Your service probably won't need to bind listeners for
	// *all* supported transports, or support both Zipkin and LightStep, and so
	// on, but we do it here for demonstration purposes.
	fs := flag.NewFlagSet("cmd", flag.ExitOnError)
	var (
		grpcAddr      = fs.String("grpc", base.GrpcAddr, "gRPC listen address")
		httpAddr      = fs.String("http", base.HttpAddr, "HTTP listen address")
		strapiAddr    = fs.String("s", "http://localhost:1337", "strapi listen address")
		siteDir       = fs.String("d", "", "Hugo site dir")
		localeDefault = fs.String("l", "en", "default locale")
		gitCommitMsg  = fs.String("m", "", "git commit message, leave blank to ignore")
	)

	fs.Usage = usageFor(fs)
	if err := fs.Parse(os.Args[1:]); err != nil {
		os.Exit(1)
	}

	// Set Strapi + Hugo site dir + git message
	impl.SetStrapiAddr(*strapiAddr)
	impl.SetSiteDir(*siteDir)
	impl.SetDefaultLocale(*localeDefault)
	impl.SetGitCommitMsg(*gitCommitMsg)

	var (
		g      run.Group
		logger log.Logger
	)

	// Create a single logger, which we'll use and give to other components.
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, base.AppName, log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	// Serve
	config := server.ServerConfig{
		GRPC: *grpcAddr,
		HTTP: *httpAddr,
	}
	server.Serve(&g, logger, config)

	// nolint:errcheck
	logger.Log("exit", g.Run())
}

func usageFor(fs *flag.FlagSet) func() {
	return func() {
		fmt.Println(base.Name, fmt.Sprintf("v%s - %s", base.Version, base.Description))
		fmt.Println("USAGE:", base.Name, "[OPTIONS] path/to/dir")
		fmt.Println("\nOPTIONS")
		fs.PrintDefaults()
	}
}
