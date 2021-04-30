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
	//account := GetGoogleAccount("ya29.a0AfH6SMDJYzoJ2z70dw-QDhQyyKGeoQfnpBkRaufEMGj25Cd7x58ZUJPqJBX8vVLX63zgYanBUb-0p6cBuC3h1npTgdu82WggyVIm7var7TwYUaLWTiuEy2YhKy1c0cUUkVNXtldb9lPWTFEnmKx_j4YB6TvhDQ")
	//println(account)
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
		ClientID:     "847931558677-h462724ru3f4ov47aennvrod20tb22rt.apps.googleusercontent.com",
		ClientSecret: "bwykyNqeTrefGoE1nvthmw6K",
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
	code := "4/0AY0e-g5B7HX2xoPu9m6InLzF5rBhkeqG32aKngJxgRIMszjv_d06jfqohRshyY-DHZYNew"
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
	println(token)
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)
	if err != nil {
		return nil
	}
	if resp == nil {
		return nil
	}
	content, err := ioutil.ReadAll(resp.Body)
	println(string(content))
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil
	}
	userInfo := &GoogleAccount{}
	err = json.Unmarshal(content, userInfo)
	if err != nil {
		return nil
	}
	return userInfo
}
