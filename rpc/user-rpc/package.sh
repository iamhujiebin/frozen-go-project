#!/usr/bin/env bash
PROJECT=user-rpc
VERSION=1.0.0

build() {
	  CGO_ENABLED=0 GOARCH=amd64 go build -o bin/user-rpc user-rpc.go || exit 1
}