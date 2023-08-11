package serviceGrpc

import (
	"github.com/atom-providers/app"
	grpcServer "github.com/atom-providers/grpc-server"
	"github.com/atom-providers/log"
	"github.com/rogeecn/atom/container"
)

func Default(providers ...container.ProviderContainer) container.Providers {
	return append(container.Providers{
		app.DefaultProvider(),
		log.DefaultProvider(),
		grpcServer.DefaultProvider(),
	}, providers...)
}
