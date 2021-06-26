package 链表

import "fmt"

var (
	ErrNoPos  = fmt.Errorf("linklist no pos")
	ErrNoHead = fmt.Errorf("linklist no head")
)

// 节点
type Node struct {
	Data interface{}
	Next *Node
}

// 链表的头节点
type Head *Node

// 链表方法(单链表/双链表/循环单链表都适合)
type ILinkList interface {
	DestroyList(head Head)                                         // 销毁链表
	InitList(...*Node) Head                                        // 初始化,一般初始化为带头节点的链表
	ClearList(head Head)                                           // 清理链表,只剩下头节点
	PrintList(head Head)                                           // 打印链表
	InsertListPos(head Head, pos int, node *Node) error            // 根据位置插入元素
	PushFront(head Head, node *Node) error                         // 插入头部
	PushBack(head Head, node *Node) error                          // 插入尾部
	DeleteNodePos(head Head, pos int) error                        // 根据位置删除元素
	PopFront(head Head) *Node                                      // 弹出第一个元素
	PopBack(head Head) *Node                                       // 弹出最后一个元素
	LengthList(head Head) int                                      // 链表长度,不算头节点
	IsEmpty(head Head) bool                                        // 是否为空链表
	LocateNodePos(head Head, pos int) *Node                        // 获取指定位置的node
	LocateNodeElem(head Head, data interface{}) *Node              // 获取指定data的node
	InsertNextNode(head Head, data interface{}, node *Node) error  // 在指定data的node后面插入node
	InsertPriorNode(head Head, data interface{}, node *Node) error // 在指定data的node前面插入node
	DeleteNode(head Head, data interface{}) error                  // 删除node
	ReverseList(head Head)                                         // 反转list
	ReversePartList(head Head, pos int) error                      // 第pos位置开始反转list
	MergeList(head1 Head, head2 Head) Head                         // 合并两个排序的list
}
