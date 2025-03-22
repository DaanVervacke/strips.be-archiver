.PHONY: build build-linux-x64 build-linux-arm build-windows

build:
	go build -o "bin/strips.be-archiver" cmd/main.go

build-linux-x64:
	GOOS=linux GOARCH=amd64 go build -o "bin/strips.be-archiver" cmd/main.go

build-linux-arm:
	GOOS=linux GOARCH=arm64 go build -o "bin/strips.be-archiver" cmd/main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o "bin/strips.be-archiver.exe" cmd/main.go

docker-build:
	docker build --tag strips.be-archiver .
