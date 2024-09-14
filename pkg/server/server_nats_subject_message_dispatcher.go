package server

import (
	"context"
	"fmt"

	nats "github.com/nats-io/nats.go"
	"github.com/qmdx00/lifecycle"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/messaging"
)

type NatsMessageDispatcher struct {
	ctx                  context.Context
	connection           messaging.NatsSubjectConnection
	listener             messaging.NatsMessageListener
	receivedMessagesChan <-chan *nats.Msg
}

func BuildNatsMessageDispatcher(connection messaging.NatsSubjectConnection, listener messaging.NatsMessageListener) Server {
	return &NatsMessageDispatcher{
		connection:           connection,
		listener:             listener,
		receivedMessagesChan: make(chan *nats.Msg),
	}
}

func (server *NatsMessageDispatcher) Run(ctx context.Context) error {

	server.ctx = ctx
	info, _ := lifecycle.FromContext(ctx)
	log.Info(fmt.Sprintf("server starting up - starting nats dispatcher %s, v.%s", info.Name(), info.Version()))

	var err error
	if _, _, server.receivedMessagesChan, err = server.connection.Connect(); err != nil {
		log.Error(fmt.Sprintf("server starting up - nats dispatcher - error: %s", err.Error()))
		return err
	}

	if err = server.ListenAndDispatch(); err != nil {
		log.Error(fmt.Sprintf("server starting up - nats dispatcher - error: %s", err.Error()))
		return err
	}

	return nil
}

func (server *NatsMessageDispatcher) ListenAndDispatch() error {

	for {
		select {
		case <-server.ctx.Done():
			return nil
		case message := <-server.receivedMessagesChan:
			go server.Dispatch(message)
		}
	}
}

func (server *NatsMessageDispatcher) Dispatch(message any) {

	var err error
	msg := message.(*nats.Msg)
	if err = server.listener.OnMessage(msg); err != nil {
		log.Error(fmt.Sprintf("nats listener - error: %s, message: %s", err.Error(), msg.Data))
	}
}

func (server *NatsMessageDispatcher) Stop(ctx context.Context) error {

	info, _ := lifecycle.FromContext(ctx)
	log.Info(fmt.Sprintf("server shutting down - stopping nats dispatcher %s, v.%s", info.Name(), info.Version()))
	server.connection.Close()
	log.Debug("server shutting down - nats dispatcher stopped")
	return nil
}
