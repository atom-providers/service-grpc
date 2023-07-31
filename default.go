package serviceGrpc

import (
	grpcServer "github.com/atom-providers/grpc-server"
	"github.com/atom-providers/log"
	"github.com/rogeecn/atom/container"
)

func Default(providers ...container.ProviderContainer) container.Providers {
	return append(container.Providers{
		log.DefaultProvider(),
		grpcServer.DefaultProvider(),
	}, providers...)
}
