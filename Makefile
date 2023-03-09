 build:
	@go build -o bin/texas-holdem

run: build
	@./bin/texas-holdem

test:
	go test -v ./... 