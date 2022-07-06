package impl

import (
	transport "github.com/go-kit/kit/transport"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"
)

func generalGrpcServerOption(logger Logger) []grpctransport.ServerOption {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}
	return options
}

func generalHttpServerOption(logger Logger) []httptransport.ServerOption {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}
	return options
}

func GrpcServerOptions(logger Logger) map[string][]grpctransport.ServerOption {
	options := map[string][]grpctransport.ServerOption{
		"entry":   entryGrpcServerOption(logger),
		"media":   mediaGrpcServerOption(logger),
		"publish": publishGrpcServerOption(logger),
	}
	return options
}

func HttpServerOptions(logger Logger) map[string][]httptransport.ServerOption {
	options := map[string][]httptransport.ServerOption{
		"entry":   entryHttpServerOption(logger),
		"media":   mediaHttpServerOption(logger),
		"publish": publishHttpServerOption(logger),
	}
	return options
}

func entryGrpcServerOption(logger Logger) []grpctransport.ServerOption {
	options := generalGrpcServerOption(logger)
	return options
}

func mediaGrpcServerOption(logger Logger) []grpctransport.ServerOption {
	options := generalGrpcServerOption(logger)
	return options
}

func publishGrpcServerOption(logger Logger) []grpctransport.ServerOption {
	options := generalGrpcServerOption(logger)
	return options
}

func entryHttpServerOption(logger Logger) []httptransport.ServerOption {
	options := generalHttpServerOption(logger)
	return options
}

func mediaHttpServerOption(logger Logger) []httptransport.ServerOption {
	options := generalHttpServerOption(logger)
	return options
}

func publishHttpServerOption(logger Logger) []httptransport.ServerOption {
	options := generalHttpServerOption(logger)
	return options
}
