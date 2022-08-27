package main

import (
	"log"
	"net"
	"net/netip"
	"os"
	"os/signal"
	"syscall"
	"yadcmd/internal/app"
	yadcmd_daemon "yadcmd/internal/pb/yadcmd.daemon"

	"google.golang.org/grpc"
)

func main() {
	const port uint16 = 49069
	listener, errMakeListener := net.Listen("tcp", net.TCPAddrFromAddrPort(
		netip.AddrPortFrom(
			netip.AddrFrom4(
				[4]byte{0, 0, 0, 0},
			),
			port,
		),
	).String())
	if errMakeListener != nil {
		panic(errMakeListener)
	}
	defer listener.Close()
	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	yadcmd_daemon.RegisterDaemonServer(server, app.NewServer())

	var errChan chan error
	go func() {
		log.Printf("gRPC server started on %d", port)
		err := server.Serve(listener)
		if err != nil {
			errChan <- err
		}
	}()

	var term chan os.Signal = make(chan os.Signal, 1)
	signal.Notify(term, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err := <-errChan:
		log.Printf("ERR:  %v", err)
		close(errChan)
		return
	case <-term:
		log.Print("INFO: done")
	}
}
