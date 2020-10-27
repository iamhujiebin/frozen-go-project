#!/usr/bin/env bash
PROJECT=base-rpc
VERSION=1.0.0

build() {
	  CGO_ENABLED=0 GOARCH=amd64 go build -o bin/base-rpc base-rpc.go || exit 1
}
