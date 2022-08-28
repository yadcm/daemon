package remote

import (
	"context"
	"log"
	"yadcmd/internal/app/config"
	"yadcmd/internal/app/listen"
	"yadcmd/internal/pb/protocol/daemon"

	"google.golang.org/grpc"
)

func Serve(ctx context.Context) error {
	conf := config.Instance.Serve.Remote
	listener, errMakeListener := listen.TCP(conf.TCP)
	if errMakeListener != nil {
		return errMakeListener
	}
	defer listener.Close()
	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	daemon.RegisterExternalAPIServer(server, newServer())
	var errChan chan error = make(chan error)
	go func() {
		log.Printf("gRPC remote started on %v", listener.Addr())
		errChan <- server.Serve(listener)
	}()

	select {
	case <-ctx.Done():
		return context.Canceled
	case err := <-errChan:
		return err
	}
}
