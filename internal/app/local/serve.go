package local

import (
	"errors"
	"log"
	"net"
	"yadcmd/internal/app/config"
	"yadcmd/internal/app/listen"
	"yadcmd/internal/pb/protocol/daemon"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var errMissingListenConfig error = errors.New("missing UNIX or TCP config")

func Serve(ctx context.Context) error {
	var listener net.Listener
	var errListen error
	if conf := config.Instance.Serve.Local.Unix; conf != nil {
		listener, errListen = listen.Unix(conf)
	} else if conf := config.Instance.Serve.Local.TCP; conf != nil {
		listener, errListen = listen.TCP(conf)
	} else {
		return errMissingListenConfig
	}
	if errListen != nil {
		return errListen
	}
	defer listener.Close()
	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	daemon.RegisterInternalAPIServer(server, newServer())
	var errChan chan error = make(chan error)
	go func() {
		log.Printf("gRPC internal started on %v", listener.Addr())
		errChan <- server.Serve(listener)
	}()

	select {
	case <-ctx.Done():
		return context.Canceled
	case err := <-errChan:
		return err
	}
}
