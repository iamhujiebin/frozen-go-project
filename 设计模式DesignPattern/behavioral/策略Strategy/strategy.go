package 策略Strategy

import "fmt"

// 策略模式
// 容器去执行策略,具体的策略可以设置,容器的入口方法不变
// 元素: 容器:Context 策略:Strategy
type Strategy interface {
	Execute()
}

type PlanAStrategy struct {
}

func (p *PlanAStrategy) Execute() {
	fmt.Println("Plan A executed")
}

type PlanBStrategy struct {
}

func (p *PlanBStrategy) Execute() {
	fmt.Println("Plan B executed")
}

func NewContext(strategy Strategy) *Context {
	return &Context{strategy: strategy}
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) Execute() {
	c.strategy.Execute()
}
