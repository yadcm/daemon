package app

import (
	"context"
	protocol "yadcmd/internal/pb/yadcmd.daemon"
)

type DaemonServer struct {
	protocol.UnimplementedDaemonServer
}

func NewServer() *DaemonServer {
	var server DaemonServer
	return &server
}

func (d *DaemonServer) AgreementStart(ctx context.Context, agreement *protocol.Hello) (*protocol.Hello, error) {
	return &protocol.Hello{}, nil
}
