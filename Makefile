APP?=hasp-exporter
VERSION?=0.0.5
COMMIT?=$(shell git rev-parse --short HEAD)
GOOS?=linux 
GOARCH?=amd64 

.PHONY: build
build:
	go build -o ${APP} -v -ldflags="-X 'main.Version=${VERSION}' -X 'main.CommitHash=${COMMIT}'" *.go

linux_build:
	GOOS=${GOOS} GOARCH=${GOARCH} make build
.DEFAULT_GOAL := build

.PHONY: rebuild
rebuild: clean build

install:
	sudo mkdir -p 	/opt/monitoring/
	sudo cp ${APP} 	/opt/monitoring/
.PHONY: clean
clean:
	@rm -f ${APP}