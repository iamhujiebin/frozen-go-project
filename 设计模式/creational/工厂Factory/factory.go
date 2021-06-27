package 工厂Factory

import "fmt"

// 针对的是单个产品等级结构
type PayChannel interface {
	CreateOrder() bool
}

type GooglePay struct {
}

func (p *GooglePay) CreateOrder() bool {
	fmt.Println("google orders ")
	return true
}

type FacebookPay struct {
}

func (p *FacebookPay) CreateOrder() bool {
	fmt.Println("facebook orders ")
	return true
}

func NewPayChannel(name string) (PayChannel, error) {
	switch name {
	case "google":
		return &GooglePay{}, nil
	case "facebook":
		return &FacebookPay{}, nil
	}
	return nil, fmt.Errorf("no pay channel:%s", name)
}
