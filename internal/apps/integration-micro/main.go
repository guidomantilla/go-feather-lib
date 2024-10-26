package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	cserver "github.com/guidomantilla/go-feather-lib/pkg/common/server"
	"github.com/guidomantilla/go-feather-lib/pkg/integration"
	"github.com/guidomantilla/go-feather-lib/pkg/integration/messaging"
)

func main() {

	_ = os.Setenv("LOG_LEVEL", "TRACE")
	cserver.Run("integration-micro", "1.0.0", func(application cserver.Application) error {

		var err error
		var headers messaging.Headers
		var message messaging.Message[string]
		var receiver messaging.ReceiverChannel[string]
		var sender messaging.SenderChannel[string]

		{
			options := messaging.HeadersOptionsChainBuilder().ErrorChannel("error-channel").ReplyChannel("reply-channel").
				Add("property", "value").Build()
			headers = messaging.NewBaseHeaders(options)
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

			message = messaging.NewBaseMessage(headers, "Hola Mundo")
			log.Info(fmt.Sprintf("Message: %v", message))

			payload := messaging.NewBaseErrorPayload("code", "message", "error")
			errMessage := messaging.NewBaseErrorMessage(headers, payload, message)
			log.Info(fmt.Sprintf("Error: %v", errMessage))
		}

		fmt.Println()
		fmt.Println()

		//

		fmt.Println()
		fmt.Println()

		{
			message = messaging.NewBaseMessage(headers, "Hola Mundo")
			receiverHandler := func(ctx context.Context, timeout time.Duration) (messaging.Message[string], error) {
				log.Debug(fmt.Sprintf("integration messaging: message arriving"))
				return message, nil
			}

			receiver = integration.BaseReceiverChannel("base-receiver-01", receiverHandler)
			message, err = receiver.Receive(context.Background(), 10*time.Second)
			log.Info(fmt.Sprintf("Done: %v, Err: %v", message, err))

			receiverHandler = func(ctx context.Context, timeout time.Duration) (messaging.Message[string], error) {
				log.Debug(fmt.Sprintf("integration messaging: message arriving"))
				<-time.After(10 * time.Second)
				return message, nil
			}
			receiver = integration.BaseReceiverChannel("base-receiver-02", receiverHandler)
			message, err = receiver.Receive(context.Background(), 5*time.Second)
			log.Info(fmt.Sprintf("Done: %v, Err: %v", message, err))
		}

		fmt.Println()
		fmt.Println()

		//

		fmt.Println()
		fmt.Println()

		{
			message = messaging.NewBaseMessage(headers, "Hola Mundo")
			senderHandler := func(ctx context.Context, timeout time.Duration, message messaging.Message[string]) error {
				log.Debug(fmt.Sprintf("integration messaging: message traveling: %v", message))
				return nil
			}

			sender = integration.BaseSenderChannel("based-sender-01", senderHandler)
			err = sender.Send(context.Background(), 10*time.Second, message)
			log.Info(fmt.Sprintf("Done: %v, Err: %v", message, err))

			senderHandler = func(ctx context.Context, timeout time.Duration, message messaging.Message[string]) error {
				<-time.After(10 * time.Second)
				log.Debug(fmt.Sprintf("integration messaging: message traveling: %v", message))
				return nil
			}

			sender = integration.BaseSenderChannel("based-sender-02", senderHandler)
			err = sender.Send(context.Background(), 5*time.Second, message)
			log.Info(fmt.Sprintf("Done: %v, Err: %v", message, err))
		}

		fmt.Println()
		fmt.Println()

		//

		fmt.Println()
		fmt.Println()

		{
			message = messaging.NewBaseMessage(headers, "Hola Mundo")
			consumerHandler := func(ctx context.Context, timeout time.Duration) (messaging.MessageStream[int], error) {
				log.Debug(fmt.Sprintf("integration messaging: message arriving: %v", message))
				inputStream := make(messaging.MessageStream[int])
				go func(inputStream messaging.MessageStream[int], headers messaging.Headers) {
					defer close(inputStream)
					for i := range 5 {
						inputStream <- messaging.NewBaseMessage(headers, i)
					}
				}(inputStream, headers)

				return inputStream, nil
			}

			var inputStream messaging.MessageStream[int]
			inputStream, err := consumerHandler(context.Background(), 10*time.Second)
			for message := range inputStream {
				log.Info(fmt.Sprintf("Done: %v, Err: %v", message, err))
			}
		}

		fmt.Println()
		fmt.Println()

		//

		fmt.Println()
		fmt.Println()

		{

		}

		return nil
	})
}
