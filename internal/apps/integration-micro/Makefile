.PHONY: phony
phony-goal: ; @echo $@

fetch-dependencies:
	go mod download

graph:
	godepgraph -s . | dot -Tpng -o godepgraph.png

imports:
	goimports-reviser -rm-unused -set-alias -format -recursive .
	go mod tidy

format:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run --max-issues-per-linter 0 --max-same-issues 0 ./...

check: fetch-dependencies graph imports format vet lint

##

update-dependencies:
	go get -u ./... && go get -t -u ./...
	go mod tidy