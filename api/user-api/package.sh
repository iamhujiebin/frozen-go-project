#!/usr/bin/env bash
PROJECT=user-api
VERSION=1.0.0

build() {
	  CGO_ENABLED=0 GOARCH=amd64 go build -o bin/user-api user.go || exit 1
}
