package local

import (
	"yadcmd/internal/app/esb"
	protocol "yadcmd/internal/pb/protocol/daemon"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type InternalServer struct {
	protocol.UnimplementedInternalAPIServer
}

func newServer() *InternalServer {
	var server InternalServer = InternalServer{}
	return &server
}

func (d *InternalServer) Setup(context.Context, *protocol.Config) (*protocol.Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Setup not implemented")
}
func (d *InternalServer) Online(info *protocol.ClientInfo, srv protocol.InternalAPI_OnlineServer) error {
	for {
		msg, err := esb.Instance.Read()
		if err != nil {
			return err
		}
		if err := srv.Send(msg); err != nil {
			return err
		}
	}
}
func (d *InternalServer) Scan(context.Context, *protocol.ScanRequest) (*protocol.ScanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Scan not implemented")
}
func (d *InternalServer) Send(context.Context, *protocol.Message) (*protocol.Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
