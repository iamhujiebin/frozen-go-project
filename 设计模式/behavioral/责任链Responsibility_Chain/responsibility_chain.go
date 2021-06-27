package 责任链Responsibility_Chain

import "fmt"

// 责任链模式
// 一个事件需要经过多个对象处理是很常见的场景
// 例子:请假！ 不同的领导能处理的请假天数不同
// 元素:请求Request 处理者Handler
// 好处就是把请假的流程入口不用改动.若流程有变,就直接改里面具体流程的就好
type Handler interface {
	HandlerRequest(days int)
	Next() Handler
}

type ConcreteHandler struct {
	Name string
	Days int
	next Handler
}

func (p *ConcreteHandler) Next() Handler {
	return p.next
}

func (p *ConcreteHandler) HandlerRequest(days int) {
	if days <= p.Days {
		fmt.Printf("%s 批准 %d 天请假\n", p.Name, days)
	} else if p.Next() != nil {
		p.Next().HandlerRequest(days)
	} else {
		fmt.Printf("没人能批准你%d天的请假", days)
	}
}

// 请假的具体handler
func NewDayOffRequest() Handler {
	C := &ConcreteHandler{
		Name: "C",
		Days: 10,
		next: nil,
	}
	B := &ConcreteHandler{
		Name: "B",
		Days: 5,
		next: C,
	}
	A := &ConcreteHandler{
		Name: "A",
		Days: 2,
		next: B,
	}
	return A
}
