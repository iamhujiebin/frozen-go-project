#!/usr/bin/env bash
PROJECT=event-rpc
VERSION=1.0.0

build() {
	  CGO_ENABLED=0 GOARCH=amd64 go build -o bin/event-rpc event-rpc.go || exit 1
}
