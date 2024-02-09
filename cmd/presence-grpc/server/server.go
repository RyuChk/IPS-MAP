package server

import (
	wiregrpc "git.cie.com/ips/wire-provider/grpc"
	"git.cie.com/ips/wire-provider/grpc/provider"
	"github.com/RyuChk/ips-map-service/cmd/presence-grpc/internal/handler"
)

type grpcServerCustomizer struct {
	handlers *handler.Handlers
}

func (gc *grpcServerCustomizer) Register(server wiregrpc.Server) error {
	handler.RegisterService(server, gc.handlers)
	return nil
}

func (gc *grpcServerCustomizer) Configrue(builder wiregrpc.Builder) error {
	builder.WithServiceName("presence-service").WithListenAddr(":6000")
	return nil
}

func ProvideGRPCServerCustomizer(handlers *handler.Handlers) provider.GRPCServerCustomizer {
	return &grpcServerCustomizer{
		handlers: handlers,
	}
}
