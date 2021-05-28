package 组合模式Composite

import "fmt"

// 组合模式
// 例子："遍历"文件夹
// 容器对象和叶子对象对外接口统一,即实现相同的interface
type Component interface {
	Traverse()
}

func NewFile(name string) *File {
	return &File{
		Name: name,
	}
}

type File struct {
	Name string
}

func (p *File) Traverse() {
	fmt.Printf("file:%s\n", p.Name)
}

func NewDirectory() *Directory {
	return &Directory{Children: make([]Component, 0)}
}

type Directory struct {
	Children []Component
}

func (p *Directory) AddComponent(c Component) {
	p.Children = append(p.Children, c)
}

func (p *Directory) Traverse() {
	for idx := range p.Children {
		p.Children[idx].Traverse()
	}
}
