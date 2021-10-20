.PHONY: run build test clean tidy download

build: clean tidy download
	go build -o bin/bot ./cmd/bot

run: tidy download
	go run ./cmd/bot

test: tidy download
	go test ./...

tidy:
	go mod tidy

download:
	go mod download

clean:
	rm -rf bin