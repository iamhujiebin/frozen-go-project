package main

import (
	"crypto/md5"
	"encoding/hex"
)

func main() {
	println(strToMd5("13660677198"))
}
func strToMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
