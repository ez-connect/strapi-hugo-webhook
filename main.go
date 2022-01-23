package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-kit/log"
	"github.com/oklog/run"

	"strapi-webhook/base"
	"strapi-webhook/base/server"
)

func main() {
	// Define our flags. Your service probably won't need to bind listeners for
	// *all* supported transports, or support both Zipkin and LightStep, and so
	// on, but we do it here for demonstration purposes.
	fs := flag.NewFlagSet("cmd", flag.ExitOnError)
	var (
		grpcAddr = fs.String("grpc", base.GrpcAddr, "gRPC listen address")
		httpAddr = fs.String("http", base.HttpAddr, "HTTP listen address")
	)

	fs.Usage = usageFor(fs)
	if err := fs.Parse(os.Args[1:]); err != nil {
		os.Exit(1)
	}

	var (
		g      run.Group
		logger log.Logger
	)

	// Create a single logger, which we'll use and give to other components.
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
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
		fmt.Println("USAGE:", base.Name, "[OPTIONS]")
		fmt.Println("\nOPTIONS")
		fs.PrintDefaults()
	}
}
