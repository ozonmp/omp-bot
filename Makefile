.PHONY: run build clean tidy download

build: clean tidy download
	go build -o bin/bot ./cmd/bot

run: tidy download
	go run ./cmd/bot

tidy:
	go mod tidy

download:
	go mod download

clean:
	rm -rf bin