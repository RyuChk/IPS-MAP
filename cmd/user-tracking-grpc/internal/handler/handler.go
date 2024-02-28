package handler

import (
	wiregrpc "git.cie-ips.com/ips/wire-provider/grpc"
	mapv1 "github.com/RyuChk/ips-map-service/internal/gen/proto/ips/map/v1"
)

type Handlers struct {
	Presence mapv1.UserTrackingServiceServer
}

func RegisterService(server wiregrpc.Server, handler *Handlers) {
	mapv1.RegisterUserTrackingServiceServer(server, handler.Presence)
}
