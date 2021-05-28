package main

import (
	"fmt"
	"net/url"
)

func main() {
	res3, _ := url.QueryUnescape("\350\256\276\350\256\241\346\250\241\345\274\217")
	fmt.Println(res3)
}
