package 链表

import (
	"fmt"
	"testing"
)

var single = SingleLinkList{}
var head Head

func init() {
	head = single.InitList(&Node{Data: 1}, &Node{Data: 2}, &Node{Data: 3})
}

func TestSingleLinkList_InitList(t *testing.T) {
	single.PrintList(head)
}

func TestSingleLinkList_LengthList(t *testing.T) {
	fmt.Printf("len:%v\n", single.LengthList(head))
}

func TestSingleLinkList_InsertListPos(t *testing.T) {
	err := single.InsertListPos(head, 3, &Node{Data: 9})
	single.PrintList(head)
	fmt.Println(err)
}

func TestSingleLinkList_PushFront(t *testing.T) {
	err := single.PushFront(head, &Node{Data: int64(11)})
	err = single.PushFront(head, &Node{Data: int64(12)})
	err = single.PushFront(head, &Node{Data: int64(13)})
	single.PrintList(head)
	fmt.Println(err)
}

func TestSingleLinkList_LocateNodeElem(t *testing.T) {
	n := single.LocateNodeElem(head, int64(11))
	fmt.Printf("node-------::%v\n", n)
}

func TestSingleLinkList_DeleteNodePos(t *testing.T) {
	err := single.DeleteNodePos(head, 1)
	single.PrintList(head)
	fmt.Println(err)
}

func TestSingleLinkList_PopFront(t *testing.T) {
	n := single.PopFront(head)
	single.PrintList(head)
	fmt.Printf("%v\n", n)
}

func TestSingleLinkList_PopBack(t *testing.T) {
	n := single.PopBack(head)
	fmt.Printf("%v\n", n)
	single.PrintList(head)
}

func TestSingleLinkList_IsEmpty(t *testing.T) {
	fmt.Printf("empty:%v\n", single.IsEmpty(head))
}

func TestSingleLinkList_LocateNodePos(t *testing.T) {
	n := single.LocateNodePos(head, 1)
	fmt.Printf("node:%v\n", n)
	n = single.LocateNodePos(head, 2)
	fmt.Printf("node:%v\n", n)
}

func TestSingleLinkList_InsertNextNode(t *testing.T) {
	err := single.InsertNextNode(head, 2, &Node{Data: 3})
	fmt.Println("err:", err)
	single.PrintList(head)
}

func TestSingleLinkList_InsertPriorNode(t *testing.T) {
	err := single.InsertPriorNode(head, 2, &Node{Data: 30})
	err = single.InsertPriorNode(head, 2, &Node{Data: 30})
	fmt.Println("err:", err)
	single.PrintList(head)
}

func TestSingleLinkList_DeleteNode(t *testing.T) {
	err := single.DeleteNode(head, 30)
	single.PrintList(head)
	fmt.Println("err:", err)
}

func TestSingleLinkList_ReverseList(t *testing.T) {
	single.PrintList(head)
	single.ReverseList(head)
	single.PrintList(head)
}

func TestSingleLinkList_ReversePartList(t *testing.T) {
	single.PrintList(head)
	err := single.ReversePartList(head, 3)
	fmt.Printf("err:%v\n", err)
	single.PrintList(head)
}

func TestSingleLinkList_MergeList(t *testing.T) {
	fmt.Println("========merge=========")
	head1 := single.InitList(&Node{Data: 1}, &Node{Data: 4}, &Node{Data: 7})
	head2 := single.InitList(&Node{Data: 2}, &Node{Data: 5}, &Node{Data: 8})
	head3 := single.InitList(&Node{Data: 3}, &Node{Data: 6}, &Node{Data: 9})
	single.PrintList(head1)
	single.PrintList(head2)
	head = single.MergeList(head1, head2)
	single.PrintList(head)
	head = single.MergeList(head3, head)
	single.PrintList(head)
}

func TestSingleLinkList_ClearList(t *testing.T) {
	single.ClearList(head)
	fmt.Printf("empty:%v\n", single.IsEmpty(head))
}
