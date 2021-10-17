.PHONY: run
run:
	go run cmd/bot/main.go

.PHONY: build
build:
	go build -o -gcflags="all=-N -l" bot cmd/bot/main.go