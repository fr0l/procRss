#!/usr/bin/env bash

set -x

rm -rf build
mkdir build

go clean
go fmt
go test -v ./...

GOOS=linux  GOARCH=amd64 go build -v -o build/procRss_linux64 procRss.go
GOOS=darwin GOARCH=amd64 go build -v -o build/procRss_mac64   procRss.go
