package tools

//go:generate mockgen -package config 		-destination ../pkg/common/config/mocks.go 		github.com/sethvargo/go-envconfig Lookuper
//go:generate mockgen -package=environment 	-destination ../pkg/common/environment/mocks.go -source ../pkg/common/environment/types.go
//go:generate mockgen -package=log 	 		-destination ../pkg/common/log/mocks.go			-source ../pkg/common/log/types.go
//go:generate mockgen -package=properties   -destination ../pkg/common/properties/mocks.go	-source ../pkg/common/properties/types.go

//go:generate mockgen -package datasource 	-destination ../pkg/datasource/mocks.go 		-source ../pkg/datasource/types.go
//go:generate mockgen -package=security 	-destination ../pkg/security/mocks.go			-source ../pkg/security/types.go
//go:generate mockgen -package=messaging 	-destination ../pkg/messaging/mocks.go 			-source ../pkg/messaging/types.go
//go:generate mockgen -package=server 		-destination ../pkg/server/mocks.go 			-source ../pkg/server/types.go 		github.com/qmdx00/lifecycle Server
