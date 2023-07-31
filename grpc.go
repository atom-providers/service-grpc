package serviceGrpc

import (
	grpcServer "github.com/atom-providers/grpc-server"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"go.uber.org/dig"
)

type GrpcService struct {
	dig.In

	Server   *grpcServer.Grpc
	Services []contracts.ServerService `group:"grpc_server_services"`
}

func Serve() error {
	return container.Container.Invoke(func(grpc GrpcService) error {
		for _, svc := range grpc.Services {
			grpc.Server.RegisterService(svc.Name(), svc.Register)
		}
		return grpc.Server.Serve()
	})
}
