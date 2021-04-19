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
	println(_FacebookConfig.AuthCodeURL(uuid.New().String()))
	ReturnAuth()
}

type FacebookAccount struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"picture"`
	Email  string `json:"email"`
	Sex    string `json:"sex"`
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
	code := "AQARaUawcRpdOeL7ZkWDED83dXCaGnYIOgXwSBjFo6DSBYL2JJMdNuKfwMrFyhDay0GHoe2cLGddHTflQk872a-V7IPKV8DFFqGnxxbMG3IB89suRJrikASMJhZqlB-ddCiA0vhciZK5CGessjMSeYc3tgXvdcIEHUTlVHBNjQOTeFmI3CkVLW8U7nmek49hPyzB8wlFoAyWblMBUqpj5uqVJo9q3oCrQQIjWnUSDowfzTdHDm9PdwpX4J6ikJSIjaAjO1fVQUB35mrfWgP-XJ1YoKEFv34uWY-stmI7G_1tiRtdXMOGnNTlofMpe4ojnzb29dzT5yphVQpUpf6oXMfy"
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
	//resp, err := http.Get("https://graph.facebook.com/me?fields=id,name,picture,hometown{location},location{location}&access_token=" + token)
	resp, err := http.Get("https://graph.facebook.com/v3.0/me?fields=id,name,picture,hometown{location{city,country,country_code}},location{location{city,country,country_code}}&access_token=" + token)
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
	err = json.Unmarshal(content, userInfo)
	println(string(content))
	if err != nil {
		return nil
	}
	return userInfo
}
