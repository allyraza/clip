
default: build

build:
	go build -o clipctl ./cmd/clipctl/main.go
	go build -o clipserver ./cmd/clipserver/main.go
