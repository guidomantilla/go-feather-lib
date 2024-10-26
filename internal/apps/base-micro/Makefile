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

check: fetch-dependencies graph imports format vet

build: check
	go build -o base-micro main.go
##

update-dependencies:
	go get -u ./... && go get -t -u ./...
	go mod tidy