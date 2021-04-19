package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
)

func main() {
	println(_GoogleConfig.AuthCodeURL(uuid.New().String()))
	ReturnAuth2()
}

type GoogleAccount struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"picture"`
	Email  string `json:"email"`
	Sex    string `json:"sex"`
}

var (
	_GoogleConfig = &oauth2.Config{
		ClientID:     "414501839443-n0vb3qb11t9f0moacc1i3hfus34rakck.apps.googleusercontent.com",
		ClientSecret: "JlS9oYJS2fspTWgAabPzE_t_",
		RedirectURL:  "https://testservice.xiangshengclub.com/api/activity/2ndcp_rank",
		//RedirectURL: "https://www.baidu.com",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
			//"https://www.googleapis.com/auth/content",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
			//AuthURL:  "http://accounts.google.inke.srv/o/oauth2/auth",
			//TokenURL: "http://accounts.google.inke.srv/o/oauth2/token",
		},
	}
)

func ReturnAuth2() {
	code := "4/0AY0e-g53rFXrLk3nUPgpUHiL5FaPrbJ2Vh7YtkE1Rq37kSJpb1zZhtz5mlNv71-x6vK1uA"
	token, err := _GoogleConfig.Exchange(context.Background(), code)
	if err != nil {
		println(err.Error())
		return
	}
	account := GetGoogleAccount(token.AccessToken)
	fmt.Printf("%#v:", account)
}

// bear token string
//如果token为空，则使用code获取token
func GetGoogleAccount(token string) *GoogleAccount {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)
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
	userInfo := &GoogleAccount{}
	println(string(content))
	err = json.Unmarshal(content, userInfo)
	if err != nil {
		return nil
	}
	return userInfo
}
