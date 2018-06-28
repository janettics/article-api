build-all: vendor build

vendor:
	go get -u github.com/bmizerany/pat
	go get -u github.com/stretchr/testify

build:
	go build ./cmd/main.go

run:
	go run ./cmd/main.go