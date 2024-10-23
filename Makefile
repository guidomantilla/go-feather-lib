.PHONY: phony
phony-goal: ; @echo $@

certificates:
	openssl genrsa -passout pass:1111 -des3 -out ca.key 4096
	openssl req -passin pass:1111 -new -x509 -days 3650 -key ca.key -out ca.crt -subj "/CN=$(SERVER_CN)"
	openssl genrsa -passout pass:1111 -des3 -out server.key 4096
	openssl req -passin pass:1111 -new -key server.key -out server.csr -subj "/CN=$(SERVER_CN)" -config $(OPENSSLCNF)
	openssl x509 -req -passin pass:1111 -days 3650 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt -extensions v3_req -extfile $(OPENSSLCNF)
	openssl pkcs8 -topk8 -nocrypt -passin pass:1111 -in server.key -out server.pem

validate: generate sort-import format vet lint coverage
	go mod tidy

generate:
	go generate ./pkg/... ./tools/...

sort-import:
	goimports-reviser -rm-unused -set-alias -format -recursive pkg
	goimports-reviser -rm-unused -set-alias -format -recursive internal

format:
	go fmt ./pkg/...

vet:
	go vet ./pkg/...

lint:
	golangci-lint run --max-issues-per-linter 0 --max-same-issues 0 ./pkg/... ./internal/...

test:
	go test -covermode count -coverprofile coverage.out.tmp.01 ./pkg/...
	cat coverage.out.tmp.01 | grep -v "mocks.go" > coverage.out
	rm coverage.out.tmp.01

coverage: test
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o .reports/coverage.html
	rm coverage.out

graph:
	godepgraph -s . | dot -Tpng -o godepgraph.png

sonarqube: coverage
	sonar-scanner

update-dependencies:
	go get -u ./...
	go get -t -u ./...
	go mod tidy

prepare:
	go install github.com/incu6us/goimports-reviser/v3@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install go.uber.org/mock/mockgen@latest
	go install github.com/cweill/gotests/gotests@latest
	go mod download
	go mod tidy