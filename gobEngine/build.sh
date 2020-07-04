#!/bin/sh

GOOS=js GOARCH=wasm go build -o main.wasm ./main.go ./resource.go ./terminal.go ./userInterface.go