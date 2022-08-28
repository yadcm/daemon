package esb

import (
	"errors"
	"yadcmd/internal/pb/protocol/daemon"
)

var Instance *esb

func init() {
	Instance = &esb{}
}

var (
	errReadChannelClosed error = errors.New("channel closed")
)

type esb struct {
	msgEvent chan *daemon.Message
}

func (e *esb) Start() {
	e.Stop()
	e.msgEvent = make(chan *daemon.Message)
}
func (e *esb) Stop() {
	close(e.msgEvent)
}

func (e *esb) Write(msg *daemon.Message) (err error) {
	defer func() {
		rec := recover()
		if rec != nil {
			err = rec.(error)
		}
	}()
	e.msgEvent <- msg
	//@todo write timeout
	return nil
}

func (e *esb) Read() (msg *daemon.Message, err error) {
	defer func() {
		rec := recover()
		if rec != nil {
			err = rec.(error)
		}
	}()
	value, opened := <-e.msgEvent
	if !opened {
		return nil, errReadChannelClosed
	}
	return value, nil
}
