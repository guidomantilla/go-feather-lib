package main

import (
	"context"
	"fmt"
	"time"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	cserver "github.com/guidomantilla/go-feather-lib/pkg/common/server"
	"github.com/guidomantilla/go-feather-lib/pkg/integration/messaging"
)

func main() {

	//_ = os.Setenv("LOG_LEVEL", "DEBUG")
	cserver.Run("nats-micro", "1.0.0", func(application cserver.Application) error {

		options := messaging.HeadersOptionsChainBuilder().ErrorChannel("error-channel").ReplyChannel("reply-channel").
			Add("property", "value").Build()
		headers := messaging.NewBasicHeaders(options)
		log.Info(fmt.Sprintf("Headers: %v", headers))

		options = messaging.NewHeadersOptions()
		headers = messaging.NewBasicHeaders(options.ErrorChannel("error-channel"), options.ReplyChannel("reply-channel"))
		log.Info(fmt.Sprintf("Headers: %v", headers))

		config := &messaging.HeadersConfig{
			ReplyChannel: "reply-channel",
			ErrorChannel: "error-channel",
		}

		options = messaging.NewHeadersOptionsFromConfig(config)
		headers = messaging.NewBasicHeaders(options)
		log.Info(fmt.Sprintf("Headers: %v", headers))

		headers = messaging.NewBasicHeadersFromConfig(config)
		log.Info(fmt.Sprintf("Headers: %v", headers))

		//

		message := messaging.NewBasicMessage(headers, "Hola Mundo")
		log.Info(fmt.Sprintf("Message: %v", message))

		payload := messaging.NewBasicErrorPayload("code", "message", "error")
		errMessage := messaging.NewBasicErrorMessage(headers, payload, message)
		log.Info(fmt.Sprintf("Error: %v", errMessage))

		//
		senderHandler := func(ctx context.Context, message messaging.Message[string], timeout time.Duration) error {
			log.Info(fmt.Sprintf("message traveling: %v", message))
			return nil
		}
		basicSender := messaging.NewBasicSenderChannel(senderHandler)
		err := basicSender.Send(context.Background(), message, 10*time.Second)
		log.Info(fmt.Sprintf("Done: %v, Err: %v", errMessage, err))

		x := messaging.NewChannelLoggerInterceptor(basicSender)
		y := messaging.NewChannelLoggerInterceptor(x)
		z := messaging.NewChannelLoggerInterceptor(y)
		err := z.Send(context.Background(), message, 10*time.Second)

		return nil
	})
}
