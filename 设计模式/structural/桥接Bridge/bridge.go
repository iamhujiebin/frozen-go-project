package 桥接Bridge

import "fmt"

// 桥接模式
// 适用于处理"多层继承结构",分类处理,然后"聚合"
// 例如: 品牌+类型,联想台式,联想平板,戴尔台式,戴尔平板
// 又例如: 画形状,颜色+图案的"聚合"
// 因为聚合,所以需要把一个类型作为另外一个类型的属性
type Shape interface {
	Draw()
	SetColor(Color)
}

type Triangle struct {
	Color Color
}

func (t *Triangle) SetColor(color Color) {
	t.Color = color
}

func (t *Triangle) Draw() {
	fmt.Printf("drawing a %s triangle\n", t.Color.Color())
}

type Rectangle struct {
	Color Color
}

func (r *Rectangle) SetColor(color Color) {
	r.Color = color
}

func (r *Rectangle) Draw() {
	fmt.Printf("drawing a %s rectangle\n", r.Color.Color())
}

type Circle struct {
	Color Color
}

func (c *Circle) SetColor(color Color) {
	c.Color = color
}

func (c *Circle) Draw() {
	fmt.Printf("drawing a %s circle\n", c.Color.Color())
}

type Color interface {
	Color() string
}

type Red struct {
}

func (r *Red) Color() string {
	return "red"
}

type Green struct {
}

func (g *Green) Color() string {
	return "green"
}

func NewShape(shape Shape, color Color) Shape {
	shape.SetColor(color)
	return shape
}
