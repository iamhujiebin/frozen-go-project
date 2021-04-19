package main

import (
	"fmt"
	"hash/fnv"
	"time"
)

func main() {
	u := time.Now().UnixNano()
	fmt.Printf("%d,%x\n", u, u)
	println(HashCode([]byte("jiebin")))
}

func HashCode(s []byte) uint32 {
	h := fnv.New32a()
	n, err := h.Write(s)
	if err != nil || n == 0 {
		return 0
	}
	return h.Sum32()
}
