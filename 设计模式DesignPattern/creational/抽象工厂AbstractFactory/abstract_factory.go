package 抽象工厂AbstractFactory

import "fmt"

// 针对的是多个产品等级结构
type Lunch interface {
	Cook()
}

type LunchFactory interface {
	CreateFood() Lunch
	CreateVegetable() Lunch
}

type Rice struct {
}

func (r *Rice) Cook() {
	fmt.Println("cooking rice")
}

type Tomato struct {
}

func (r *Tomato) Cook() {
	fmt.Println("cooking tomato")
}

type SimpleLunch struct {
}

func (s *SimpleLunch) CreateFood() Lunch {
	return &Rice{}
}

func (s *SimpleLunch) CreateVegetable() Lunch {
	return &Tomato{}
}

func NewSimpleLunch() LunchFactory {
	return &SimpleLunch{}
}
