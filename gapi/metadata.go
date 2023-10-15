package gapi

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	// gateway
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	xForwardedForHeader        = "x-forwarded-for"
	// grpc
	userAgentHeader = "user-agent"
)

type Metadata struct {
	UserAgent string
	ClientIp  string
}

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	data := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if useragents := md.Get(grpcGatewayUserAgentHeader); len(useragents) > 0 {
			data.UserAgent = useragents[0]
		}

		if useragents := md.Get(userAgentHeader); len(useragents) > 0 {
			data.UserAgent = useragents[0]
		}

		if clientIps := md.Get(xForwardedForHeader); len(clientIps) > 0 {
			data.ClientIp = clientIps[0]
		}
	}

	if pr, ok := peer.FromContext(ctx); ok {
		data.ClientIp = pr.Addr.String()
	}

	return data
}
