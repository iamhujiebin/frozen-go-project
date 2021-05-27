package 适配器Adapter

import "fmt"

// 适配器模式
// 目标方法跟调用者的方法不一致,导致没法执行
// 三个元素:target adaptee adapter
// adaptee 想要执行Execute,但只有SpecialExecute,需要adapter帮手
// adapter需要"继承"adaptee,并且实现Target
type Target interface {
	Execute()
}

type Adaptee struct {
}

func (p *Adaptee) SpecialExecute() {
	fmt.Println("Adaptee SpecialExecute")
}

type Adapter struct {
	*Adaptee
}

func (a *Adapter) Execute() {
	a.SpecialExecute()
}

func NewAdapter(adaptee *Adaptee) *Adapter {
	return &Adapter{adaptee}
}
