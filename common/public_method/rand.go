package public_method

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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

// 获取指定长度的随机字符串
func GetRandomString(n int) string {
	base := primitive.NewObjectID().Hex()
	baseBytes := []byte(base)
	result := []byte{}
	for i := 0; i < n; i++ {
		result = append(result, baseBytes[rand.Intn(len(baseBytes))])
	}
	return string(result)
}
