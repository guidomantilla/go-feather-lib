module go-feather-lib/internal/apps/rabbitmq-micro

go 1.23.1

require (
	github.com/guidomantilla/go-feather-lib v0.0.0
	github.com/rabbitmq/amqp091-go v1.10.0
)

replace github.com/guidomantilla/go-feather-lib v0.0.0 => ../../../

require (
	github.com/avast/retry-go/v4 v4.6.0 // indirect
	github.com/qmdx00/lifecycle v1.1.1 // indirect
	go.uber.org/mock v0.5.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
)
