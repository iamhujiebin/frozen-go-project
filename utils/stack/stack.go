package stack

import "fmt"

var (
	ErrNoInit = fmt.Errorf("no stack init")
	ErrNode   = fmt.Errorf("not a node")
	ErrFull   = fmt.Errorf("stack is full")
)

type Node struct {
	Data interface{}
	Next *Node
}

type Stack struct {
	top *Node // 栈顶数据
	Len int   // 栈元素个数
	Cap int   // 栈总容量
}

type IStack interface {
	InitStack(cap int, nodes ...*Node) (*Stack, error) // 初始化栈
	DestroyStack(*Stack)                               // 销毁栈
	ClearStack(*Stack)                                 // 清空栈
	Push(*Stack, *Node) error                          // 压入栈
	Pop(*Stack) *Node                                  // 弹出栈
	IsEmpty(*Stack) bool                               // 是否为空
	IsFull(*Stack) bool                                // 是否满了
	Length(*Stack) int                                 // 栈的长度
	GetTop(*Stack) *Node                               // 获取栈顶元素
	PrintStack(*Stack)                                 // 打印Stack的所有元素
}
