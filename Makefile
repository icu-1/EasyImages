BUILD_ENV := CGO_ENABLED=0
PROJECT := $(shell go list -m)
VERSION := $(shell cat VERSION)
DIR := $(shell pwd)
LDFLAGS := -ldflags "-w -s -X ${PROJECT}/vars.Version=${VERSION} -X ${PROJECT}/vars.Project=${PROJECT} -X ${PROJECT}/vars.ProjectDir=${DIR}"
TARGET_EXEC := server

.PHONY: all clean setup build-linux build-osx build-windows copy

all: clean setup build-linux build-osx build-windows copy

clean:
	rm -rf bin

setup:
	mkdir -p bin/linux
	mkdir -p bin/osx
	mkdir -p bin/windows

copy: clean setup
	cp config.yaml bin/config.yaml

build-linux: copy
	${BUILD_ENV} GOARCH=amd64 GOOS=linux go build ${LDFLAGS} -o bin/linux/${TARGET_EXEC} -trimpath cmd/command.go

build-osx: copy
	${BUILD_ENV} GOARCH=amd64 GOOS=darwin go build ${LDFLAGS} -o bin/osx/${TARGET_EXEC} -trimpath cmd/command.go

build-windows: copy
	${BUILD_ENV} GOARCH=amd64 GOOS=windows go build ${LDFLAGS} -o bin/windows/${TARGET_EXEC}.exe -trimpath cmd/command.go

