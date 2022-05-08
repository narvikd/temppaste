#!/bin/sh
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w"
docker build -t temppaste .
docker save -o temppaste.tar temppaste:latest
