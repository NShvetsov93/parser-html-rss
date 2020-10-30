.PHONY: build
build:
	go build -o ./bin ./cmd/server

.PHONY: gen
gen:
	protoc --go_out=plugins=grpc:pkg/ -I api/ api/api.proto

.PHONY: evans
	evans api/api.proto -p 8080
