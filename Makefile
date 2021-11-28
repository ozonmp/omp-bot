.PHONY: run
run:
	go run cmd/bot/main.go

.PHONY: build
build:
	go build -o bot cmd/bot/main.go

# ----------------

.PHONY: generate
generate:
	protoc --proto_path=protos protos/ozonmp/cnm_film_api/v1/*.proto --go_out=plugins=grpc:pb

.PHONY: remove-pb
generate:
	rm -rf pb/*
