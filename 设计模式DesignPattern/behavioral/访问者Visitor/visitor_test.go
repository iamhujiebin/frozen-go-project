package 访问者Visitor

import "testing"

func TestVisitor(t *testing.T) {
	chrome := &Chrome{}
	firefox := &Firefox{}
	douyin := &Douyin{}
	weibo := &Weibo{}
	chrome.Accept(douyin)
	firefox.Accept(douyin)
	chrome.Accept(weibo)
	firefox.Accept(weibo)
}
