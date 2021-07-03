package 二叉树

import (
	"fmt"
	"testing"
)

var bst *BNode
var bs BinarySortTree

func init() {
	bst = bs.InitTree(5, 2, 3, 4, 1, 6, 7, 8, 9, 0)
}

func TestBinarySortTree_InitTree(t *testing.T) {
	fmt.Println(bs.LevelOrder(bst))
}

func TestBinarySortTree_LevelOrder(t *testing.T) {
	fmt.Println(bs.LevelOrder(bst))
}

func TestBinarySortTree_PreOrder(t *testing.T) {
	res := make([]interface{}, 0)
	bs.PreOrder(bst, &res)
	fmt.Println(res)
}

func TestBinarySortTree_InOrder(t *testing.T) {
	res := make([]interface{}, 0)
	bs.InOrder(bst, &res)
	fmt.Println(res)
}

func TestBinarySortTree_PostOrder(t *testing.T) {
	res := make([]interface{}, 0)
	bs.PostOrder(bst, &res)
	fmt.Println(res)
}

func TestBinarySortTree_TreeDepth(t *testing.T) {
	fmt.Println(bs.TreeDepth(bst))
}
