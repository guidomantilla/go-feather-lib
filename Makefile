.PHONY: phony
phony-goal: ; @echo $@

certificates:
	# Inspired from: https://github.com/grpc/grpc-java/tree/master/examples#generating-self-signed-certificates-for-use-with-grpc

    # Output files
    # ca.key: Certificate Authority private key file (this shouldn't be shared in real-life)
    # ca.crt: Certificate Authority trust certificate (this should be shared with users in real-life)
    # server.key: Server private key, password protected (this shouldn't be shared)
    # server.csr: Server certificate signing request (this should be shared with the CA owner)
    # server.crt: Server certificate signed by the CA (this would be sent back by the CA owner) - keep on server
    # server.pem: Conversion of server.key into a format gRPC likes (this shouldn't be shared)

    # Step 1: Generate Certificate Authority + Trust Certificate (ca.crt)
	openssl genrsa -passout pass:1111 -des3 -out ca.key 4096
	openssl req -passin pass:1111 -new -x509 -days 3650 -key ca.key -out ca.crt -subj "/CN=$(SERVER_CN)"

	# Step 2: Generate the Server Private Key (server.key)
	openssl genrsa -passout pass:1111 -des3 -out server.key 4096

	# Step 3: Get a certificate signing request from the CA (server.csr)
	openssl req -passin pass:1111 -new -key server.key -out server.csr -subj "/CN=$(SERVER_CN)" -config $(OPENSSLCNF)

	# Step 4: Sign the certificate with the CA we created (it's called self signing) - server.crt
	openssl x509 -req -passin pass:1111 -days 3650 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt -extensions v3_req -extfile $(OPENSSLCNF)

	# Step 5: Convert the server certificate to .pem format (server.pem) - usable by gRPC
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
	golangci-lint run ./pkg/...

test:
	go test -covermode count -coverprofile coverage.out.tmp.01 ./pkg/...
	cat coverage.out.tmp.01 | grep -v "mocks.go" > coverage.out

coverage: test
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

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