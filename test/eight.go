package main

import (
	"flag"
	"fmt"
	"net/url"
	"strings"
)

//\350\256\276\350\256\241\346\250\241\345\274\217
func main() {
	s := flag.String("s", "", "string")
	flag.Parse()
	str := strings.Replace(*s, "\\\\", "\\", -1)
	res3, _ := url.QueryUnescape(str)
	fmt.Println(res3)
}
