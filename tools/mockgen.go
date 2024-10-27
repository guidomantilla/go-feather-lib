package tools

//go:generate mockgen -package=config 		-destination ../pkg/common/config/mocks.go 				github.com/sethvargo/go-envconfig Lookuper
//go:generate mockgen -package=environment 	-destination ../pkg/common/environment/mocks.go 		-source ../pkg/common/environment/types.go
//go:generate mockgen -package=log 	 		-destination ../pkg/common/log/mocks.go					-source ../pkg/common/log/types.go
//go:generate mockgen -package=properties   -destination ../pkg/common/properties/mocks.go			-source ../pkg/common/properties/types.go

//go:generate mockgen -package=gorm 		-destination ../pkg/datasource/gorm/mocks.go 			-source ../pkg/datasource/gorm/types.go
//go:generate mockgen -package=mongo 		-destination ../pkg/datasource/mongo/mocks.go 			-source ../pkg/datasource/mongo/types.go
//go:generate mockgen -package=security 	-destination ../pkg/security/authentication_mocks.go	-source ../pkg/security/authentication_types.go
//go:generate mockgen -package=security 	-destination ../pkg/security/authorization_mocks.go		-source ../pkg/security/authorization_types.go
//go:generate mockgen -package=security 	-destination ../pkg/security/password_mocks.go			-source ../pkg/security/password_types.go
//go:generate mockgen -package=security 	-destination ../pkg/security/principal_manager_mocks.go -source ../pkg/security/principal_manager_types.go
//go:generate mockgen -package=security 	-destination ../pkg/security/token_manager_mocks.go 	-source ../pkg/security/token_manager_types.go

//go:generate mockgen -package=amqp 		-destination ../pkg/messaging/rabbitmq/amqp/mocks.go 	-source ../pkg/messaging/rabbitmq/amqp/types.go
//go:generate mockgen -package=streams 		-destination ../pkg/messaging/rabbitmq/streams/mocks.go -source ../pkg/messaging/rabbitmq/streams/types.go

//go:generate mockgen -package=server 		-destination ../pkg/server/mocks.go 					-source ../pkg/server/types.go
