package 原型Prototype

// 原型模式
// 克隆新的结构体
type Prototype interface {
	Name() string
	Clone() Prototype
}

func NewConcretePrototype() *ConcretePrototype {
	c := new(ConcretePrototype)
	c.name = "concrete prototype"
	return c
}

type ConcretePrototype struct {
	name string
}

func (p *ConcretePrototype) Name() string {
	return p.name
}

func (p *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{
		name: p.name,
	}
}
