#!/bin/bash

rm -r dist/
env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./dist/main ./cmd/main.go
cd ./dist/ && zip main.zip main && cd ..
cp ./dist/main.zip ./cmd/