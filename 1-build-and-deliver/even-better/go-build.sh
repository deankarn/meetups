#!/bin/bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -ldflags "-s" -installsuffix cgo -o app-linux-amd64 .

