package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Unix()
	println(now)
	println(now >> 1)
	println(now >> 2)
	println(now >> 7)
	println(1 << 32)
	idStr := fmt.Sprintf("%d%06d", now, 1%1000000)
	println(idStr)
	println(now - 1430409600)
}
