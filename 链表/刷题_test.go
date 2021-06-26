package 链表

import (
	"fmt"
	"testing"
)

var head2 Head
var roundHead Head
var single2 = SingleLinkList{}

var head3, head4 Head

func init() {
	head2 = single.InitList(&Node{Data: 1}, &Node{Data: 2},
		&Node{Data: 3}, &Node{Data: 4}, &Node{Data: 5}, &Node{Data: 6}, &Node{Data: 7}, &Node{Data: 8})
	// 造个环
	roundHead = single.InitList(&Node{Data: 1}, &Node{Data: 2}, &Node{Data: 3})
	cur := roundHead
	var n *Node
	for cur.Next != nil {
		if cur.Data == 2 {
			n = cur
		}
		cur = cur.Next
	}
	cur.Next = n

	n1, n2, n3, n4, n5 := &Node{Data: 1}, &Node{Data: 2}, &Node{Data: 3}, &Node{Data: 4}, &Node{Data: 5}
	head3 = single2.InitList(n1, n3, n4, n5)
	n2.Next = n3
	head4 = single2.InitList(n2)
}

func TestPrintReverse(t *testing.T) {
	single2.PrintList(head2)
	PrintReverse2(head2)
	//PrintReverse(head2)
}

func TestFindLastK(t *testing.T) {
	single2.PrintList(head2)
	fmt.Println(FindLastK(head2, 1))
	fmt.Println(FindLastK(head2, 8))
	fmt.Println(FindLastK(head2, 3))
	fmt.Println(FindLastK(head2, 9))
}

func TestFindRoundEntry(t *testing.T) {
	fmt.Println(FindRoundEntry(head2))
	fmt.Println(FindRoundEntry2(head2))
	fmt.Println(FindRoundEntry(roundHead))
	fmt.Println(FindRoundEntry2(roundHead))
}

func TestFindCrossEntry(t *testing.T) {
	single2.PrintList(head2)
	single2.PrintList(head3)
	single2.PrintList(head4)
	fmt.Println(FindCrossEntry(head2, head3))
	fmt.Println(FindCrossEntry2(head2, head3))
	fmt.Println(FindCrossEntry(head3, head4))
	fmt.Println(FindCrossEntry2(head3, head4))
}

func TestHeadWithTail(t *testing.T) {
	single2.PrintList(head2)
	HeadWithTail(head2)
	single2.PrintList(head2)
}

func TestHeadWithTail2(t *testing.T) {
	single2.PrintList(head2)
	HeadWithTail2(head2)
	single2.PrintList(head2)
}
