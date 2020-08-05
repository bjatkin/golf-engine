#!/bin/sh
go generate
GOOS=darwin GOARCH=amd64 go build -o mac/golf_toolkit *.go
GOOS=linux GOARCH=amd64 go build -o linux/golf_toolkit *.go
GOOS=windows GOARCH=amd64 go build -o windows/golf_toolkit *.go