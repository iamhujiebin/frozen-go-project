package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/demo/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("前缀匹配"))
		return
	})
	http.HandleFunc("/demo1", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("精准匹配"))
		return
	})
	err := http.ListenAndServe(":6666", nil)
	panic(err)

	//e := gin.Default()
	////e.GET("/ping", func(context *gin.Context) {})
	////e.GET("/pong", func(context *gin.Context) {})
	////e.GET("/user/:id", func(context *gin.Context) {})
	////e.GET("/user/:id/add", func(context *gin.Context) {})
	////err := e.Run(":7777")
	//panic(err)
}
