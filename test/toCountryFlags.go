package main


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"unicode"
)

func main() {
	type Data = struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	flagArr := make([]Data, 0)
	file, err := os.Open("./flag.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(b, &flagArr)
	file2, err := os.Create("target.txt")
	if err != nil {
		panic(err)
	}
	for _, v := range flagArr {
		country := v.Name
		country = strings.Replace(v.Name, "flag_", "", 1)
		country = strings.Replace(country, "_", "", -1)
		country = strings.Replace(country, ".png", "", -1)
		country = Ucfirst(country)
		text := fmt.Sprintf("[country.%s]\nflag = \"%s\"\n", country, v.Url)
		file2.WriteString(text)
	}
	file2.Close()
}

func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func uploadFile(filePath string) (uploadPath string) {
	u := "http://upload.hkainob.com/upload/image?sufix=png"
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	res, err := http.Post(u, "binary/octet-stream", file)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	type uploadRes struct {
		DmError  int64  `json:"dm_error"`
		ErrorMsg string `json:"error_msg"`
		Url      string `json:"url"`
	}

	var resp uploadRes
	err = json.Unmarshal(body, &resp)
	if err != nil {
		panic(err)
		return
	}
	if resp.DmError != 0 {
		return
	}
	uploadPath = resp.Url
	return
}
