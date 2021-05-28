package 享元FlyWeight

import "sync"

// 共享变量池！
// 获取时候,有就返回,没有就创建
type FlyWeight struct {
	Name string
}

func NewFlyWeight(name string) *FlyWeight {
	return &FlyWeight{
		Name: name,
	}
}

type FlyWeightFactory struct {
	pool map[string]*FlyWeight
	mux  sync.Mutex
}

func (p *FlyWeightFactory) GetFlyWeight(name string) *FlyWeight {
	p.mux.Lock()
	defer p.mux.Unlock()
	if _, ok := p.pool[name]; ok {
		return p.pool[name]
	}
	p.pool[name] = NewFlyWeight(name)
	return p.pool[name]
}

func NewFlyWeightFactory() *FlyWeightFactory {
	return &FlyWeightFactory{
		pool: make(map[string]*FlyWeight),
	}
}
