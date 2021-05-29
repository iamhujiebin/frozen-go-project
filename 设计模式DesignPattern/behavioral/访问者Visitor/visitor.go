package 访问者Visitor

import "fmt"

// 访问者模式
// 把对象和行为分开来抽象成"对象",对象中有行为,行为中有对象
// 元素: 对象Element 行为Visit
// Element 有accept,visitor有visit
// 例子:爬虫,用不同的浏览器访问不同的网站

type IVisitor interface {
	Name() string
	Visit(element IElement)
}

type IElement interface {
	Name() string
	Accept(visitor IVisitor)
}

type Firefox struct {
}

func (c *Firefox) Name() string {
	return "firefox"
}

func (c *Firefox) Accept(visitor IVisitor) {
	visitor.Visit(c)
}

type Chrome struct {
}

func (c *Chrome) Name() string {
	return "chrome"
}

func (c *Chrome) Accept(visitor IVisitor) {
	visitor.Visit(c)
}

type Weibo struct {
}

func (*Weibo) Name() string {
	return "weibo"
}

func (w *Weibo) Visit(element IElement) {
	fmt.Printf("use %s visit %s\n", element.Name(), w.Name())
}

type Douyin struct {
}

func (*Douyin) Name() string {
	return "douyin"
}

func (w *Douyin) Visit(element IElement) {
	fmt.Printf("use %s visit %s\n", element.Name(), w.Name())
}
