package tools

//go:generate mockgen -package config 		-destination ../pkg/config/mocks.go 		github.com/sethvargo/go-envconfig Lookuper
//go:generate mockgen -package datasource 	-destination ../pkg/datasource/mocks.go 	-source ../pkg/datasource/types.go
//go:generate mockgen -package=environment 	-destination ../pkg/environment/mocks.go 	-source ../pkg/environment/types.go
//go:generate mockgen -package=log 	 		-destination ../pkg/log/mocks.go			-source ../pkg/log/types.go
