.PHONY: phony
phony-goal: ; @echo $@

install: fetch-dependencies
	go install github.com/incu6us/goimports-reviser/v3@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install go.uber.org/mock/mockgen@latest
	go install github.com/vladopajic/go-test-coverage/v2@latest
	npm install @jacksontian/gocov -g

fetch-dependencies:
	go mod download

generate:
	go generate ./pkg/... ./tools/...

graph:
	godepgraph -s ./pkg/datasource/gocql | dot -Tpng -o ./pkg/datasource/gocql/gocql.png
	godepgraph -s ./pkg/datasource/goredis | dot -Tpng -o ./pkg/datasource/goredis/goredis.png
	godepgraph -s ./pkg/datasource/gorm | dot -Tpng -o ./pkg/datasource/gorm/gorm.png
	godepgraph -s ./pkg/datasource/mongo | dot -Tpng -o ./pkg/datasource/mongo/mongo.png
	godepgraph -s ./pkg/integration | dot -Tpng -o ./pkg/integration/integration.png
	godepgraph -s ./pkg/integration/messaging | dot -Tpng -o ./pkg/integration/messaging/messaging.png
	godepgraph -s ./pkg/messaging/rabbitmq/amqp | dot -Tpng -o ./pkg/messaging/rabbitmq/amqp/amqp.png
	godepgraph -s ./pkg/messaging/rabbitmq/streams | dot -Tpng -o ./pkg/messaging/rabbitmq/streams/streams.png
	godepgraph -s ./pkg/security | dot -Tpng -o ./pkg/security/security.png
	godepgraph -s ./pkg/server | dot -Tpng -o ./pkg/server/server.png

imports:
	goimports-reviser -rm-unused -set-alias -format -recursive pkg
	goimports-reviser -rm-unused -set-alias -format -recursive internal
	go mod tidy

format:
	go fmt ./pkg/...

vet:
	go vet ./pkg/...

lint:
	golangci-lint run --max-issues-per-linter 0 --max-same-issues 0 ./pkg/... ./internal/...

test:
	go test -covermode atomic -coverprofile .reports/coverage.out.tmp ./pkg/...
	cat .reports/coverage.out.tmp | grep -v "mocks.go" > .reports/coverage.out && rm .reports/coverage.out.tmp

coverage-report: test
	gocov .reports/coverage.out && rm -R .reports/coverage || true  && cp -R coverage .reports && rm -R coverage

coverage-check: coverage-report
	go-test-coverage --config=.testcoverage.yml

check: fetch-dependencies generate graph imports format vet lint test

validate: check coverage-check

##

prepare: install
	go install github.com/cweill/gotests/gotests@latest

update-dependencies:
	go get -u ./... && go get -t -u ./...
	go mod tidy

certificates:
	openssl genrsa -passout pass:1111 -des3 -out ca.key 4096
	openssl req -passin pass:1111 -new -x509 -days 3650 -key ca.key -out ca.crt -subj "/CN=$(SERVER_CN)"
	openssl genrsa -passout pass:1111 -des3 -out server.key 4096
	openssl req -passin pass:1111 -new -key server.key -out server.csr -subj "/CN=$(SERVER_CN)" -config $(OPENSSLCNF)
	openssl x509 -req -passin pass:1111 -days 3650 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt -extensions v3_req -extfile $(OPENSSLCNF)
	openssl pkcs8 -topk8 -nocrypt -passin pass:1111 -in server.key -out server.pem
