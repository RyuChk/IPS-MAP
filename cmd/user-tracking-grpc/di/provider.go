package di

import (
	grpcProvide "git.cie.com/ips/wire-provider/grpc/provider"
	"github.com/RyuChk/ips-map-service/cmd/user-tracking-grpc/internal/handler"
	"github.com/RyuChk/ips-map-service/cmd/user-tracking-grpc/server"
	mapv1 "github.com/RyuChk/ips-map-service/internal/gen/proto/ips/map/v1"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	server.ProvideGRPCServerCustomizer,
	handler.ProvideUserTrackingServer,
	wire.Struct(new(handler.Handlers), "*"),
)

type Locator struct {
	Handler              *handler.Handlers
	GRPCServerCustomizer grpcProvide.GRPCServerCustomizer
	Map                  mapv1.MapServiceServer
}
