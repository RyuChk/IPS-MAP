package di

import wiregrpc "git.cie-ips.com/ips/wire-provider/grpc"

type Container struct {
	server wiregrpc.Server
}
