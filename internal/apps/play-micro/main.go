package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	cserver "github.com/guidomantilla/go-feather-lib/pkg/common/server"
	"github.com/guidomantilla/go-feather-lib/pkg/integration/messaging"
)

func main() {

	_ = os.Setenv("LOG_LEVEL", "DEBUG")
	cserver.Run("nats-micro", "1.0.0", func(application cserver.Application) error {

		options := messaging.HeadersOptionsChainBuilder().ErrorChannel("error-channel").ReplyChannel("reply-channel").
			Add("property", "value").Build()
		headers := messaging.NewBaseHeaders(options)
		log.Info(fmt.Sprintf("Headers: %+v", headers))

		options = messaging.NewHeadersOptions()
		headers = messaging.NewBaseHeaders(options.ErrorChannel("error-channel"), options.ReplyChannel("reply-channel"))
		log.Info(fmt.Sprintf("Headers: %v", headers))

		config := &messaging.HeadersConfig{
			ReplyChannel: "reply-channel",
			ErrorChannel: "error-channel",
		}

		options = messaging.NewHeadersOptionsFromConfig(config)
		headers = messaging.NewBaseHeaders(options)
		log.Info(fmt.Sprintf("Headers: %v", headers))

		headers = messaging.NewBaseHeadersFromConfig(config)
		log.Info(fmt.Sprintf("Headers: %v", headers))

		//

		message := messaging.NewBaseMessage(headers, "Hola Mundo")
		log.Info(fmt.Sprintf("Message: %v", message))

		payload := messaging.NewBaseErrorPayload("code", "message", "error")
		errMessage := messaging.NewBaseErrorMessage(headers, payload, message)
		log.Info(fmt.Sprintf("Error: %v", errMessage))

		//
		senderHandler := func(ctx context.Context, message messaging.Message[string], timeout time.Duration) error {
			log.Debug(fmt.Sprintf("integration messaging: message traveling: %v", message))
			return nil
		}

		var sender messaging.SenderChannel[string]

		sender = messaging.NewLoggedSenderChannel("logged-sender-01", senderHandler)
		err := sender.Send(context.Background(), message, 10*time.Second)
		log.Info(fmt.Sprintf("Done: %v, Err: %v", errMessage, err))

		senderHandler = func(ctx context.Context, message messaging.Message[string], timeout time.Duration) error {
			<-time.After(10 * time.Second)
			log.Debug(fmt.Sprintf("integration messaging: message traveling: %v", message))
			return nil
		}
		sender = messaging.NewLoggedSenderChannel("logged-sender-02", senderHandler)
		sender = messaging.NewTimeoutSenderChannel("timeout-sender-02", sender)
		err = sender.Send(context.Background(), message, 5*time.Second)
		log.Info(fmt.Sprintf("Done: %v, Err: %v", errMessage, err))

		return nil
	})
}
