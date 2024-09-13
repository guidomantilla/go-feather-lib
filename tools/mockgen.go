package tools

//go:generate mockgen -package config 		-destination ../pkg/config/mocks.go 		github.com/sethvargo/go-envconfig Lookuper
//go:generate mockgen -package datasource 	-destination ../pkg/datasource/mocks.go 	-source ../pkg/datasource/types.go
//go:generate mockgen -package=environment 	-destination ../pkg/environment/mocks.go 	-source ../pkg/environment/types.go
//go:generate mockgen -package=log 	 		-destination ../pkg/log/mocks.go			-source ../pkg/log/types.go
//go:generate mockgen -package=log 	 		-destination ../pkg/log/mocks.go			-source ../pkg/log/types.go
//go:generate mockgen -package=properties   -destination ../pkg/properties/mocks.go		-source ../pkg/properties/types.go
//go:generate mockgen -package=security 	-destination ../pkg/security/mocks.go		-source ../pkg/security/types.go
//go:generate mockgen -package=server 		-destination ../pkg/server/mocks.go 		github.com/qmdx00/lifecycle Server
