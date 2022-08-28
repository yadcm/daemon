package remote

import (
	"context"
	"yadcmd/internal/app/esb"
	protocol "yadcmd/internal/pb/protocol/daemon"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ExternalServer struct {
	protocol.UnimplementedExternalAPIServer
}

func newServer() *ExternalServer {
	var server ExternalServer
	return &server
}

func (e *ExternalServer) Info(context.Context, *protocol.Host) (*protocol.Host, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (e *ExternalServer) OnMessage(ctx context.Context, msg *protocol.Message) (*protocol.Message, error) {
	go esb.Instance.Write(msg)
	//@todo implement storage
	return nil, status.Errorf(codes.Unimplemented, "method OnMessage not implemented")
}
