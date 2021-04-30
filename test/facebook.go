package main

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
)

func main() {
	//println(_FacebookConfig.AuthCodeURL(uuid.New().String()))
	//ReturnAuth()
	account := GetFacebookAccount("e")
	println(account)
}

type FacebookAccount struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Picture struct {
		Data struct {
			Height       int    `json:"height"`
			Width        int    `json:"width"`
			Url          string `json:"url"`
			IsSilhouette bool   `json:"is_silhouette"`
		} `json:"data"`
	} `json:"picture"`
}

var (
	_FacebookConfig = &oauth2.Config{
		ClientID:     "826021178006005",
		ClientSecret: "a9fafcc5dfb62b0481128eae23c3483c",
		RedirectURL:  "https://testservice.xiangshengclub.com/api/activity/2ndcp_rank",
		Scopes:       []string{"public_profile", "email", "user_hometown", "user_location"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.facebook.com/dialog/oauth",
			TokenURL: "https://graph.facebook.com/oauth/access_token",
			//AuthURL:  "http://www.facebook.inke.srv/dialog/oauth",
			//TokenURL: "http://www.facebook.inke.srv/oauth/access_token",
		},
	}
)

func ReturnAuth() {
	code := "AQBz0LeIm0VdTAvzIz28ZBJBGFgvWG1e24bfZ7A7c6Jqt2Flrjrq_7eZu0PgnmQHEgUnA9nu43jXgc917RPcJp2lG3p5642rd7KQ7fGCo1iH1lMkhqAKasxWsLvAle0TAm_PbpqtmlVn1y5O18uDUFmTnrmwVOeVniL3yOHRAVndsZ3SQbxCOwp9DgYmqfQVNcKUgpgtVGOaLS9uq6F4Zg1HmbyvUtlMTWzf2KpL7W1UwIT9P5jI4EO_2I6DGOSn7vNGqn43AZ3mZJrYus8AFoKIX20oVRcXeK10R0EFu75h8KsZTagxYIlkYglTmKk1vOZthQIK0BvIClUMVtE7-wR1"
	token, err := _FacebookConfig.Exchange(context.Background(), code)
	if err != nil {
		println(err.Error())
		return
	}
	//p.Token = token.AccessToken
	account := GetFacebookAccount(token.AccessToken)
	if account == nil {
		return
	}
	fmt.Printf("%#v\n", account)
}

// bear token string
func GetFacebookAccount(token string) *FacebookAccount {
	if token == "" {
		return nil
	}
	token = "EAAEOCGhGaL0BALs4AkZBJ2pZBTMPwtMf0tv0ljgkZBbqSRq90GC3fA5hxCT1bndwPvQ5fGPh3XmlMZAQCNjJcgCZAGWerO80YzFZAkSPxySvTMtxQ0OlgHL2iYXOV5eWziFTyB8I1he8FDXtyUwvZCBbMtZAp5vSL2EVjSgbUS14YZABdkI881wDbKvNhZCNa4vSyNNotzrLc3WW9NeBzWUfufBLNUtZAAUSZBgZD"
	println(token)
	resp, err := http.Get("https://graph.facebook.com/me?fields=id,name,picture&access_token=" + token)
	//resp, err := http.Get("https://graph.facebook.com/me?fields=id,name,picture,hometown{location{city,country,country_code}},location{location{city,country,country_code}}&access_token=" + token)
	if err != nil {
		return nil
	}
	if resp == nil {
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	userInfo := new(FacebookAccount)
	err = json.Unmarshal(content, &userInfo)
	println(string(content))
	if err != nil {
		return nil
	}
	return userInfo
}

func UploadFile(fburl string) (uploadPath string, err error) {
	resp, err := http.Get("http://graph.facebook.com/1865661313609928/picture")
	url := fmt.Sprintf("https://upload.meetstarlive.com/upload/media?sufix=png")
	res, err := http.Post(url, "binary/octet-stream", resp.Body)
	if err != nil {
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}
	return
}
