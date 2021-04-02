package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"sync"

	"github.com/gogf/gf/frame/g"
)

type repoInfo struct {
	Url  string `json:"ssh_url_to_repo"`
	Name string `json:"path_with_namespace"`
}

type groupInfo struct {
	Id int `json:"id"`
}

const (
	goPath = "/Users/jiebin/go/src/code.inke.cn/"
	//gitlabToken = "L7-VN6FNDkx3u4oe2PsH" // gitlabToken获取 从gitlab setting中的Access Tokens获取
	gitlabToken = "4DSM3BFxTvdXzZyt7Uaj" // gitlabToken获取 从gitlab setting中的Access Tokens获取
	gitlabAddr  = "code.inke.cn"
)

var (
	target = []string{
		// "ss",
		// "cj",
		// "ms",
		// "so",
		//"deep",
		// "seek",
		"yochat",
	}
	// 音泡代号为cj  香芋味ms 77为ss  填满为拉取整个广州服务端后端代码
	ssGroupUrl     = "https://code.inke.cn/api/v4/groups/1888/subgroups"
	cjGroupUrl     = "https://code.inke.cn/api/v4/groups/1243/subgroups"
	msGroupUrl     = ""
	soGroupUrl     = "https://code.inke.cn/api/v4/groups/2828/subgroups"
	deepGroupUrl   = "https://code.inke.cn/api/v4/groups/2819/subgroups"
	seekGroupUrl   = "https://code.inke.cn/api/v4/groups/3165/subgroups"
	yochatGroupUrl = "https://code.inke.cn/api/v4/groups/3459/subgroups"
)

func getGroupId(url string) []*groupInfo {
	queryUrl := fmt.Sprintf("%s?private_token=%s", url, gitlabToken)
	content := g.Client().GetBytes(queryUrl)
	fmt.Println(string(content))
	var groupInfoList []*groupInfo
	err := json.Unmarshal(content, &groupInfoList)
	if err != nil {
		fmt.Println(err)
	}
	return groupInfoList
}

func getNext(groupId int) {
	url := genNextUrl(groupId)
	content := g.Client().GetBytes(url)
	var repoInfoList []*repoInfo
	err := json.Unmarshal(content, &repoInfoList)
	if err != nil {
		fmt.Println(err)
	}
	wg := sync.WaitGroup{}
	for _, v := range repoInfoList {
		fmt.Println(v.Url)
		wg.Add(1)
		go func(v *repoInfo) {
			defer wg.Done()
			var (
				command string
				out     bytes.Buffer
				stderr  bytes.Buffer
			)
			path := goPath + v.Name
			fmt.Printf("path:%v\n", path)
			_, err := os.Stat(path)
			if err != nil {
				command = fmt.Sprintf("git clone %s %s", v.Url, path)
			} else {
				command = fmt.Sprintf("git -C \"%s\" pull", path)
			}
			cmd := exec.Command("/bin/bash", "-c", command)
			cmd.Stdout = &out
			cmd.Stderr = &stderr
			if err = cmd.Run(); err != nil {
				fmt.Println("==========")
				fmt.Println(err, stderr.String())
				fmt.Println("拉取", v.Name, "失败")
				fmt.Println()
				fmt.Println("==========")
				return
			}
			// fmt.Println("正在拉取 ", v.Name, "...")
			fmt.Println(out.String())
		}(v)
	}
	wg.Wait()
}

func genNextUrl(groupId int) string {
	return fmt.Sprintf("https://%s/api/v4/groups/%d/projects?per_page=100&private_token=%s", gitlabAddr, groupId, gitlabToken)
}

func main() {
	for _, code := range target {
		switch code {
		case "cj":
			for _, group := range getGroupId(cjGroupUrl) {
				getNext(group.Id)
			}
		case "ms":
			for _, group := range getGroupId(msGroupUrl) {
				getNext(group.Id)
			}
		case "ss":
			for _, group := range getGroupId(ssGroupUrl) {
				getNext(group.Id)
			}
		case "so":
			for _, group := range getGroupId(soGroupUrl) {
				getNext(group.Id)
			}
		case "deep":
			for _, group := range getGroupId(deepGroupUrl) {
				getNext(group.Id)
			}
		case "seek":
			for _, group := range getGroupId(seekGroupUrl) {
				getNext(group.Id)
			}
		case "yochat":
			for _, group := range getGroupId(yochatGroupUrl) {
				getNext(group.Id)
			}
		}
	}
}
