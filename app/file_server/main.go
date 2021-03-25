package main

import (
	"frozen-go-project/app/file_server/core"
	"frozen-go-project/app/file_server/initialize"
)

func main() {
	router := initialize.Routers()
	address := ":8888"
	s := core.InitServer(address, router)
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
