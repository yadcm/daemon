package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"yadcmd/internal/app/local"
	"yadcmd/internal/app/remote"
)

func main() {
	var errChan chan error = make(chan error)
	var wg sync.WaitGroup
	wg.Add(2)
	ctxRemote, cancelRemote := context.WithCancel(context.Background())
	ctxLocal, cancelLocal := context.WithCancel(context.Background())
	go run(ctxRemote, remote.Serve, &wg, errChan)
	go run(ctxLocal, local.Serve, &wg, errChan)

	var term chan os.Signal = make(chan os.Signal, 1)
	signal.Notify(term, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err := <-errChan:
		log.Printf("err: %v", err)
		cancelRemote()
		cancelLocal()
		wg.Wait()
		close(errChan)
		return
	case <-term:
		log.Println("shutting down")
		cancelRemote()
		cancelLocal()
		wg.Wait()
		return
	}
}

type serveFn func(context.Context) error

func run(ctx context.Context, serve serveFn, wg *sync.WaitGroup, errChan chan error) {
	err := serve(ctx)
	wg.Done()
	if err != nil {
		errChan <- err
	}
}
