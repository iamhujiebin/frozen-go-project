#!/usr/bin/env bash
PROJECT=event-worker
VERSION=1.0.0

build() {
	  CGO_ENABLED=0 GOARCH=amd64 go build -o bin/event-worker event-worker.go || exit 1
}
