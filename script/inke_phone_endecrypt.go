package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

type PhoneSecret struct {
	Phone       string
	SecretPhone string
}

const sourceAesKey = "xiangyuxingqiu_user_account_base"

func main() {
	str, _ := NewPhoneSecret("8612345678170")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678171")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678172")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678173")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678174")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678175")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678176")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678177")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678178")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678179")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678180")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678181")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678182")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678183")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678184")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678185")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678185")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678186")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678187")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678188")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678189")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8613660677198")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8611223344556")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8613268076059")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8613660677199")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8613660677190")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8613660677191")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678900")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678901")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678902")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678903")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678904")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678905")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8613345678900")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8611112222333")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678911")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8618682145110")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678169")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678168")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678167")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678166")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678165")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678164")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678163")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678162")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678161")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678160")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678191")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678192")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678193")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8613660677197")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678100")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678101")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678102")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678103")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678104")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678105")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678106")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678107")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678108")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678109")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8612345678110")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8615360603975")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("8615118077687")
	fmt.Printf("\"%s\",", str.SecretPhone)
	str, _ = NewPhoneSecret("85261595593")
	fmt.Printf("\"%s\",", str.SecretPhone)
}

func NewPhoneSecret(phone string) (*PhoneSecret, error) {
	isPlain, err := isPlainPhone(phone)
	if err != nil {
		return nil, err
	}

	s := &PhoneSecret{}
	if isPlain {
		s.Phone = phone
		var bphone []byte = []byte(phone)
		var bkey []byte = []byte(sourceAesKey)
		srca := EncryptAES(bphone, bkey)
		hexsrc := hex.EncodeToString(srca)
		s.SecretPhone = hexsrc
		return s, nil
	}

	hexdst := []byte{}
	hexdst, _ = hex.DecodeString(phone)
	var bkey []byte = []byte(sourceAesKey)
	desa := DecryptAES(hexdst, bkey)
	s.Phone = string(desa[:])
	return s, nil
}

func isPlainPhone(phone string) (bool, error) {
	isPlain := true
	for _, ch := range phone {
		if ch >= 'a' && ch < 'g' {
			isPlain = false
			continue
		}

		if ch >= '0' && ch <= '9' {
			continue
		}

		return isPlain, fmt.Errorf("Phone %v is invalid", phone)
	}
	return isPlain, nil
}

func EncryptAES(src, key []byte) []byte {
	//1.创建并返回一个使用DES算法的cipher.Block接口。
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//2.对src进行填充
	src = padding(src, block.BlockSize())
	//3.返回blockModel
	//vi := []byte("aaaabbbb")
	//blockModel := cipher.NewCBCEncrypter(block, vi)
	//str,_ = "key[:block.BlockSize()]", key[:block.BlockSize()])
	blockModel := cipher.NewCBCEncrypter(block, key[:block.BlockSize()]) //block.BlockSize() ==len(key)
	//4.crypto加密连续块
	blockModel.CryptBlocks(src, src)

	return src
}

func padding(src []byte, blockSize int) []byte {
	//func padding(src []byte, blockSize int) {
	//1.截取加密代码 段数
	padding := blockSize - len(src)%blockSize
	//2.有余数
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	//3.添加余数
	src = append(src, padText...)
	return src

}

func Depadding(src []byte) []byte {
	//1.取出最后一个元素
	lasteum := int(src[len(src)-1])
	//2.删除和最后一个元素相等长的字节
	//str,_ = "src", src)
	newText := src[:len(src)-lasteum]
	return newText
}

//aes解密
func DecryptAES(src, key []byte) []byte {
	//1.创建并返回一个使用DES算法的cipher.Block接口。
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//2.crypto解密
	//vi := []byte("aaaabbbb")
	//str,_ = "src[:block.BlockSize()]", key[:block.BlockSize()])
	blockModel := cipher.NewCBCDecrypter(block, key[:block.BlockSize()]) //block.BlockSize() ==len(key)
	//3.解密连续块
	blockModel.CryptBlocks(src, src)
	//.删除填充数组
	src = Depadding(src)

	return src
}
