BUILD_ENV := CGO_ENABLED=0
BUILD=`date +%FT%T%z`
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD}"

TARGET_EXEC := server

.PHONY: all clean setup build-linux build-osx build-windows copy

all: clean setup build-linux build-osx build-windows copy

clean:
	rm -rf bin

setup:
	mkdir -p bin/linux
	mkdir -p bin/osx
	mkdir -p bin/windows

build-linux: setup
	${BUILD_ENV} GOARCH=amd64 GOOS=linux go build ${LDFLAGS} -o bin/linux/${TARGET_EXEC} -trimpath cmd/command.go

build-osx: setup
	${BUILD_ENV} GOARCH=amd64 GOOS=darwin go build ${LDFLAGS} -o bin/osx/${TARGET_EXEC} -trimpath cmd/command.go

build-windows: setup
	${BUILD_ENV} GOARCH=amd64 GOOS=windows go build ${LDFLAGS} -o bin/windows/${TARGET_EXEC}.exe -trimpath cmd/command.go

copy: setup
	cp config.yaml bin/config.yaml