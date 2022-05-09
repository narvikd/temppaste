#!/bin/sh
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w"
docker build -t narvikd/temppaste . --no-cache
