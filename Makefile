.PHONY: help
help:
	@echo 'help for help'
	@echo ''

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: test
test: tidy
	go test ./...
	golangci-lint run --max-issues-per-linter 0 --max-same-issues 0

.PHONY: build-windows
build-windows: test
	GOOS=windows GOARCH=amd64 go build -o sharpey-windows.exe ./cmd/sharpey/sharpey.go 

.PHONY: build-linux
build-linux: test
	GOOS=linux GOARCH=amd64 go build -o sharpey-linux ./cmd/sharpey/sharpey.go 

.PHONY: build-all
build-all: build-windows build-linux


proto:
	protoc --doc_out=./docs --doc_opt=markdown,protobuf.md --proto_path ./api/grpc/proto/ --go_opt module=github.com/supanygga/sharpey --go_out module=github.com/supanygga/sharpey:./ --go-grpc_out module=github.com/supanygga/sharpey:./ ./api/grpc/proto/*.proto
