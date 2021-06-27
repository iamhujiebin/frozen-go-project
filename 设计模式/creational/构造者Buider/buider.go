package 构造者Buider

// 构造者模式,目标是构造出"产品"
// 有builder、director(指挥者)、product
type Builder interface {
	BuildPart1()
	BuildPart2()
	GetResult() *Product
}

type Director struct {
	builder Builder
}

func (p *Director) Construct() *Product {
	p.builder.BuildPart1()
	p.builder.BuildPart2()
	return p.builder.GetResult()
}

type Product struct {
	Part1 string
	Part2 string
}

func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{}
}

type ConcreteBuilder struct {
	part1 string
	part2 string
}

func (c *ConcreteBuilder) BuildPart1() {
	c.part1 = "build part1"
}

func (c *ConcreteBuilder) BuildPart2() {
	c.part2 = "build part2"
}

func (p *ConcreteBuilder) GetResult() *Product {
	return &Product{
		Part1: p.part1,
		Part2: p.part2,
	}
}

func NewDirector(b Builder) *Director {
	return &Director{builder: b}
}
