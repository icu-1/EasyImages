#!/bin/bash

outdir="bin"
rm -rf ${outdir}
GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o ${outdir}/linux-server -trimpath cmd/command.go
cp config.yaml $outdir/config.yaml