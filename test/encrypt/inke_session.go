package main

import (
	"crypto/md5"
	"crypto/rc4"
	"encoding/base64"
	"fmt"
	"github.com/prometheus/common/log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	USER_NAME_REG = "{$user_name}"
	GIFT_NAME_REG = "{$gift_name}"
)

func main() {
	data := fmt.Sprintf("%s send gift %s", USER_NAME_REG, GIFT_NAME_REG)
	rep := strings.Replace(data, USER_NAME_REG, "jiebin", -1)
	rep = strings.Replace(rep, GIFT_NAME_REG, "mengyin", -1)
	fmt.Println(rep)
	fmt.Println(fmt.Sprintf("%v%02d", "user_relation_", 191305%1))
	fmt.Println(191305 % 1)

	fmt.Println(sessionDecode(191305, "20DVyjY8YVovKXNHBi0VaVTnU0jeF5MYPuaki0BZArli21LtRwjQi3"))

	mp := make(map[ClientType]string)
	l := []interface{}{}
	uid := uint64(191305)
	reply := []string{
		"206Zl3scLKJA6trpjRN71DYQptwi0N4EU8121H5USXi1qBh1l9wi3",
		"20jFIT1Kvv6rEwi1DjT8q1jRJW5mpCCi0OnlbZDusdQe0i13r3zsi3",
		"20DVyjY8YVovKXNHBi0VaVTnU0jeF5MYPuaki0BZArli21LtRwjQi3",
	}
	for i := range reply {
		isExp, _, clientType := sessionDecode(uid, reply[i])
		if isExp {
			l = append(l, reply[i])
			continue
		}
		val, ex := mp[clientType]
		if ex {
			l = append(l, val)
		}
		mp[clientType] = reply[i]
	}
	fmt.Println(mp, l)
}

type ClientType int32

const (
	CLIENT_APP ClientType = iota
	CLIENT_WEB
)
const (
	SID_VERSION = "20"
	RC4_PASS    = "Wdi2vloq"
	LETTERS     = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func sessionDecode(uid uint64, sid string) (bool, uint64, ClientType) {
	isExp := true
	expAt := uint64(0)

	clientType := CLIENT_APP
	if uid == 0 || sid == "" {
		log.Infof("uid %v or sid %v is invalid", uid, sid)
		return isExp, expAt, clientType
	}

	if sid[0:2] != SID_VERSION {
		log.Infof("sid (%v) head of uid (%v) check failed", sid, uid)
		return isExp, expAt, clientType
	}

	session := sid[2:]
	salt := session[4:12]
	body := session[0:4] + session[12:]
	dst, err := b64decode(body)
	if err != nil {
		return isExp, expAt, clientType
	}

	key := getRC4Key(salt)
	src, _ := rc4c(dst, key)
	log.Debugf("rc4 key-%v src-%v", key, src)

	// uid#salt#time#salt#client_type
	l := strings.Split(src, "#")
	if len(l) < 4 {
		return isExp, expAt, clientType
	}

	iuid, _ := strconv.ParseUint(l[0], 10, 64)
	if uid != iuid {
		return isExp, expAt, clientType
	}

	expAt, err = strconv.ParseUint(l[2], 10, 64)

	if expAt > uint64(time.Now().Unix()) {
		isExp = false
	}

	if len(l) == 5 && l[4] == strconv.Itoa(int(CLIENT_WEB)) {
		clientType = CLIENT_WEB
	}
	log.Debugf("decode sid-%v isExp-%v expAt-%v clientType-%v", sid, isExp, expAt, clientType)
	return isExp, expAt, clientType
}

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = LETTERS[rand.Intn(len(LETTERS))]
	}
	return string(b)
}

func getRC4Key(salt string) string {
	// hasher := md5.New()
	// hasher.Write([]byte(salt + RC4_PASS))
	// return hex.EncodeToString(hasher.Sum(nil))
	return fmt.Sprintf("%x", md5.Sum([]byte(salt+RC4_PASS)))
}

func rc4c(src, key string) (string, error) {
	c, err := rc4.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	dst := make([]byte, len(src))
	c.XORKeyStream(dst, []byte(src))
	return string(dst), nil
}

func b64encode(src string) string {
	s := base64.StdEncoding.EncodeToString([]byte(src))
	var dst string
	// for i := range s {
	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "i":
			dst = dst + "i0"
		case "+":
			dst = dst + "i1"
		case "/":
			dst = dst + "i2"
		case "=":
			dst = dst + "i3"
		default:
			dst = dst + string(s[i])
		}
	}

	return dst
}

func b64decode(src string) (string, error) {
	var dst string
	// for i := range src {
	for i := 0; i < len(src); i++ {
		if string(src[i]) == "i" {
			i++
			if i < len(src) {
				switch string(src[i]) {
				case "0":
					dst = dst + "i"
				case "1":
					dst = dst + "+"
				case "2":
					dst = dst + "/"
				case "3":
					dst = dst + "="
				default:
					dst = dst + "i" + string(src[i])
				}
			} else {
				dst = dst + "i"
			}

		} else {
			dst = dst + string(src[i])
		}
	}
	data, err := base64.StdEncoding.DecodeString(dst)
	if err != nil {
		log.Infof("decode base64 dst:%s error: %v", dst, err)
		return string(data), err
	}
	log.Debugf("dst (%v) data (%s)", dst, data)

	return string(data), nil
}
