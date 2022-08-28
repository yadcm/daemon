package remote

import (
	"context"
	"net/netip"
	"time"

	protocol "yadcmd/internal/pb/protocol/daemon"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	readTimeout    time.Duration = 5 * time.Second
	writeTimeout   time.Duration = 5 * time.Second
	connectTimeout time.Duration = 2 * time.Second
)

type Client struct {
	protocol.ExternalAPIClient
	conn *grpc.ClientConn
}

func NewClient(addr netip.AddrPort) (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()
	conn, err := grpc.DialContext(ctx,
		addr.String(),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		return nil, err
	}
	var cl Client
	cl.ExternalAPIClient = protocol.NewExternalAPIClient(conn)
	cl.conn = conn
	return &cl, nil
}

func (c *Client) Close() {
	c.conn.Close()
}
