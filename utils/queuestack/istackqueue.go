package queuestack

import "fmt"

var (
	ErrNoInit = fmt.Errorf("no init")
	ErrFull   = fmt.Errorf("is full")
)

type node struct {
	Data interface{}
	Next *node

	priority int // 优先队列-优先级
}

// 队列|栈
type QueueStack struct {
	Len int // 栈|队列元素个数
	Cap int // 栈|队列总容量 ps:链栈可以不需要,顺序栈扩容可以用到

	top *node // 栈顶数据-链栈

	arr      []interface{} // 节点数组-顺序栈|循环队列
	topIndex int           // top节点下标-顺序栈

	headIndex int // 头节点索引-循环队列
	tailIndex int // 尾节点索引-循环队列

	head *node // 头节点-链式队列|优先队列
	tail *node // 尾节点-链式队列
}

type IQueueStack interface {
	Init(cap int, datas ...interface{}) (*QueueStack, error) // 初始化栈|队列
	Destroy(*QueueStack)                                     // 销毁栈|队列
	Clear(*QueueStack)                                       // 清空栈|队列
	Push(*QueueStack, interface{}) error                     // 压入栈|队列
	Pop(*QueueStack) interface{}                             // 弹出栈|队列
	IsEmpty(*QueueStack) bool                                // 是否为空
	IsFull(*QueueStack) bool                                 // 是否满了
	Length(*QueueStack) int                                  // 栈|队列的长度
	GetTop(*QueueStack) interface{}                          // 获取栈|队列顶元素
	Print(*QueueStack)                                       // 打印QueueStack的所有元素
}
