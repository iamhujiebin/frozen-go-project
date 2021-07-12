package 链表

import (
	"fmt"
	"testing"
)

func TestNewSkipList(t *testing.T) {
	skipList := NewSkipList(4)
	skipList.Insert(1, 1)
	skipList.Insert(2, 2)
	skipList.Insert(3, 3)
	skipList.Insert(4, 3)
	skipList.Insert(5, 3)
	skipList.Insert(6, 3)
	skipList.Insert(7, 3)
	skipList.Insert(8, 3)
	skipList.Insert(9, 3)
	skipList.Insert(10, 3)
	n := skipList.Search(3)
	fmt.Printf("n:%+v\n", n)
	skipList.Delete(3)
	n = skipList.Search(3)
	fmt.Printf("n:%+v\n", n)
	n = skipList.Search(5)
	fmt.Printf("n:%+v\n", n)
	n = skipList.Search(6)
	fmt.Printf("n:%+v\n", n)
	n = skipList.Search(11)
	fmt.Printf("n:%+v\n", n)
}
