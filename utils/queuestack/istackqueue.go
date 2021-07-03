package queuestack

import "fmt"

var (
	ErrNoInit = fmt.Errorf("no init")
	ErrNode   = fmt.Errorf("not a node")
	ErrFull   = fmt.Errorf("is full")
)

// todo 这里的Node,不应该暴露出来
// 应该要把Data interface{}作为外面就行了
type Node struct {
	Data interface{}
	Next *Node
}

// 队列|栈
type QueueStack struct {
	Len int // 栈|队列元素个数
	Cap int // 栈|队列总容量 ps:链栈可以不需要,顺序栈扩容可以用到

	top *Node // 栈顶数据-链栈

	arr      []*Node // 节点数组-顺序栈|循环队列
	topIndex int     // top节点下标-顺序栈

	headIndex int // 头节点索引-循环队列
	tailIndex int // 尾节点索引-循环队列

	head *Node // 头节点-链式队列
	tail *Node // 尾节点-链式队列
}

type IQueueStack interface {
	Init(cap int, nodes ...*Node) (*QueueStack, error) // 初始化栈|队列
	Destroy(*QueueStack)                               // 销毁栈|队列
	Clear(*QueueStack)                                 // 清空栈|队列
	Push(*QueueStack, *Node) error                     // 压入栈|队列
	Pop(*QueueStack) *Node                             // 弹出栈|队列
	IsEmpty(*QueueStack) bool                          // 是否为空
	IsFull(*QueueStack) bool                           // 是否满了
	Length(*QueueStack) int                            // 栈|队列的长度
	GetTop(*QueueStack) *Node                          // 获取栈|队列顶元素
	Print(*QueueStack)                                 // 打印QueueStack的所有元素
}
