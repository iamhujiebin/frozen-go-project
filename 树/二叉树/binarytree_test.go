package 二叉树

import (
	"fmt"
	"testing"
)

var binaryTree Root
var bt BinaryTree

func init() {
	binaryTree = bt.InitTree(nil)
}

func TestBinaryTree_TreeDepth(t *testing.T) {
	t.Log(bt.TreeDepth(binaryTree))
}

func TestBinaryTree_LevelOrder(t *testing.T) {
	fmt.Printf("levelOrder:%v\n", bt.LevelOrder(binaryTree))
}
