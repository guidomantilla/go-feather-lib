package messaging

import (
	"fmt"
	"strings"
	"time"

	nats "github.com/nats-io/nats.go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type DefaultNatsSubjectConnection struct {
	url                        string
	connection                 *nats.Conn
	notifyOnClosedConnection   chan any //TODO:
	subscription               *nats.Subscription
	notifyOnClosedSubscription chan any //TODO:
	subject                    string
	queue                      string
	isReady                    bool
	notifyOnClosingWatcher     chan bool
	receivedMessagesChan       chan *nats.Msg
}

func NewDefaultNatsSubjectConnection(url string, username string, password string, subject string, queue string) *DefaultNatsSubjectConnection {

	if strings.TrimSpace(url) == "" {
		log.Fatal("starting up - error setting up nats connection: url is empty")
	}

	if strings.TrimSpace(username) == "" {
		log.Fatal("starting up - error setting up nats connection: username is empty")
	}

	if strings.TrimSpace(password) == "" {
		log.Fatal("starting up - error setting up nats connection: password is empty")
	}

	if strings.TrimSpace(subject) == "" {
		log.Fatal("starting up - error setting up nats connection: subject is empty")
	}

	if strings.TrimSpace(queue) == "" {
		log.Fatal("starting up - error setting up nats connection: queue is empty")
	}

	url = strings.Replace(url, ":username", username, 1)
	url = strings.Replace(url, ":password", password, 1)

	return &DefaultNatsSubjectConnection{
		url:                    url,
		subject:                subject,
		queue:                  queue,
		isReady:                false,
		notifyOnClosingWatcher: make(chan bool),
		receivedMessagesChan:   make(chan *nats.Msg),
	}
}

//

func (client *DefaultNatsSubjectConnection) Start() {
	go client.watchConnection()
	<-time.After(time.Second)
}

func (client *DefaultNatsSubjectConnection) Close() {
	client.notifyOnClosingWatcher <- true
}

func (client *DefaultNatsSubjectConnection) Connect() (*nats.Conn, *nats.Subscription, chan *nats.Msg, error) {

	if client.isReady {
		return client.connection, client.subscription, client.receivedMessagesChan, nil
	}

	<-time.After(makeConnectionDelay)
	if !client.isReady {
		return nil, nil, nil, fmt.Errorf("unable to connect to %s", client.url)
	}

	return client.connection, client.subscription, client.receivedMessagesChan, nil
}

//

func (client *DefaultNatsSubjectConnection) makeConnection() error {

	var err error
	if client.connection, err = nats.Connect(client.url); err != nil {
		return err
	}

	if client.subscription, err = client.connection.ChanQueueSubscribe(client.subject, client.queue, client.receivedMessagesChan); err != nil {
		return err
	}

	client.notifyOnClosedConnection = make(chan any, 1)
	//client.connection.ClosedHandler()(client.connection)

	client.notifyOnClosedSubscription = make(chan any, 1)
	//client.subscription.SetClosedHandler()

	client.isReady = true
	return nil
}

func (client *DefaultNatsSubjectConnection) watchConnection() {

	log.Info("nats - connection demon started")
	defer log.Info("nats - connection demon stopped")

	for {

		if !client.isReady {
			var err error
			if err = client.makeConnection(); err != nil {
				log.Debug("nats - connection demon - failed to connect. retrying...")
				continue
			}
			log.Info("nats - connection ready")
		}

		select {

		case <-client.notifyOnClosedConnection:
			client.isReady = false
			log.Info("nats - connection closed. reconnecting...")
			continue

		case <-client.notifyOnClosedSubscription:
			client.isReady = false
			log.Info("nats - subscription closed. recreating...")
			continue

		case <-time.After(5 * time.Second):
			if client.isReady {
				log.Info("nats - mock ping...")
				/*
					var err error
					if _, err = client.channel.QueueInspect(client.queue.Name); err != nil { //nolint:staticcheck
						client.isReady = false
						log.Info("nats - failed to ping the queue. notifying...")
						continue
					}
				*/
			}
			continue

		case <-client.notifyOnClosingWatcher:
			return
		}
	}
}
