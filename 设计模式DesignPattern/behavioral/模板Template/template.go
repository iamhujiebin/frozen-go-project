package 模板Template

import "fmt"

// 模版模型
// 定义好一个模版,具体的实现都会遵循定义的模版去走
// 例如下单流程
type CreateOrderTemplate interface {
	PreCheck()
	CreateOrder()
	PostCheck()
}

func NewWorker(w CreateOrderTemplate) *CreateOrderWorker {
	return &CreateOrderWorker{w}
}

type CreateOrderWorker struct {
	CreateOrderTemplate // 这样子写不太好理解,其实这个是需要初始化的,不然会报空指针
	//t CreateOrderTemplate // 这样子写好理解一点
}

func (p *CreateOrderWorker) Create() {
	p.PreCheck()
	p.CreateOrder()
	p.PostCheck()
}

type GoogleWorker struct {
}

func (g *GoogleWorker) PreCheck() {
	fmt.Println("google pre check")
}

func (g *GoogleWorker) CreateOrder() {
	fmt.Println("google create order")
}

func (g *GoogleWorker) PostCheck() {
	fmt.Println("google post check")
}
