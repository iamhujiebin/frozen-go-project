package main

import (
	"fmt"
	"net/url"
)

func main() {
	picture, err := url.Parse("https://platform-lookaside.fbsbx.com/platform/profilepic/?asid=1865661313609928&height=50&width=50&ext=1622197680&hash=AeQhi_2TDasa2BrLSrw")
	if err == nil {
		query := picture.Query()
		query.Set("width", "500")
		query.Set("height", "500")
		picture.RawQuery = query.Encode()
	}
	println(picture.String())
	fmt.Printf("mac input test")
	// this is english
	// this is english
	// english for input
}
