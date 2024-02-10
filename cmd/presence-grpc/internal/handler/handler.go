package handler

import (
	wiregrpc "git.cie.com/ips/wire-provider/grpc"
	mapv1 "github.com/RyuChk/ips-map-service/internal/gen/proto/ips/map/v1"
)

type Handlers struct {
	Presence mapv1.PresenceServiceServer
}

func RegisterService(server wiregrpc.Server, handler *Handlers) {
	mapv1.RegisterPresenceServiceServer(server, handler.Presence)
}
