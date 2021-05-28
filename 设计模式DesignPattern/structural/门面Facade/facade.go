package 门面Facade

import "fmt"

// 门面模式
// 提供一个统一入口,屏蔽里面具体的子系统
type Facade struct {
	Light        *Light
	TV           *TV
	AirCondition *AirCondition
}

func NewLight() *Light {
	return &Light{}
}

type Light struct {
}

func (p *Light) Open() {
	fmt.Println("light open")
}

func NewTv() *TV {
	return &TV{}
}

type TV struct {
}

func (p *TV) Open() {
	fmt.Println("tv open")
}

func NewAirCondition() *AirCondition {
	return &AirCondition{}
}

type AirCondition struct {
}

func (p *AirCondition) Open() {
	fmt.Println("air-condition open")
}

func (p *Facade) Open() {
	p.Light.Open()
	p.TV.Open()
	p.AirCondition.Open()
}

func NewFacade() *Facade {
	return &Facade{
		Light:        NewLight(),
		TV:           NewTv(),
		AirCondition: NewAirCondition(),
	}
}
