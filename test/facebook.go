package main

import (
	"context"
	"encoding/json"
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
		Scopes:       []string{"public_profile", "email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.facebook.com/v4.0/dialog/oauth",
			TokenURL: "https://graph.facebook.com/v4.0/oauth/access_token",
		},
	}
)

func ReturnAuth() {
	code := "AQAXXUSF66Qj0didoB_u6IHGDDaujM5dnQvx6ktAhZZFF3clg4rIvwscuCqZjOyxnw8vRK04rRQx-5S5nkOeIwC4Y6QULP_pVxzaHhFAgK1Rk7z9PId0YluQdn0PmIAXHgT8Q2s0J6pBDq2YG7SJnyBQbRKbLSJtnk9DgRxiocvSLdfTCjkxwBoxGPOGiU8xKWJDqobMvEWwQIeilPtHwobzXabs8FHC69heBbARcV2w6ZHhLaNRIZbXq_MzTYTXWnZ6yRTyhFUPE7o2eWDsg6Jhxf2DMl3RfL199Tm4nHY3BXLfJwCqfhHK09NT5mheXrGXnrPJMS0WeZFkauKKVxGR"
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
	println(account)
}

// bear token string
func GetFacebookAccount(token string) *FacebookAccount {
	if token == "" {
		return nil
	}
	resp, err := http.Get("https://graph.facebook.com/v4.0/me?access_token=" + token)
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
	if err != nil {
		return nil
	}
	userInfo.Avatar = "https://graph.facebook.com/" + userInfo.ID + "/picture?width=9999"
	return userInfo
}
