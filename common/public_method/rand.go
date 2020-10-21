package public_method

import (
	"math/rand"
	"time"
)

//获取n个随机的0-9的纯数字字符
func GetRandomCodeFromNumber(n int) string {
	letters := []rune("0123456789")
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rd.Intn(len(letters))]
	}
	return string(b)
}
