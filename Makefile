all: install

install: go.sum
	go install ./cmd/server

go.sum: go.mod
	go mod verify