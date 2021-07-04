package 二叉树

import (
	. "frozen-go-project/utils/queuestack"
)

type BinaryTree struct {
}

/*
	 1
   /  \
  2    3
 /  \   /
 4   5 6
/ \  /\
7  8 9 0
*/
func (B BinaryTree) InitTree(datas ...interface{}) *bNode {
	//if len(datas) <= 0 {
	//	return nil
	//}
	n1, n2, n3, n4, n5, n6, n7, n8, n9, n10 :=
		&bNode{data: 1}, &bNode{data: 2}, &bNode{data: 3}, &bNode{data: 4}, &bNode{data: 5},
		&bNode{data: 6}, &bNode{data: 7}, &bNode{data: 8}, &bNode{data: 9}, &bNode{data: 0}
	n1.left, n1.right = n2, n3
	n2.left, n2.right = n4, n5
	n3.left = n6
	n4.left, n4.right = n7, n8
	n5.left, n5.right = n9, n10
	return n1
}

// 递归查找
// 递归查找貌似没发找到pre指针
func (B BinaryTree) Find(root *bNode, data interface{}) (node *bNode, pre *bNode) {
	if root == nil {
		return
	}
	if root.data == data {
		node = root
		return
	}
	node, pre = B.Find(root.left, data)
	if node != nil {
		return
	}
	return B.Find(root.right, data)
}

func (B BinaryTree) Insert(root *bNode, data interface{}) *bNode {
	panic("implement me")
}

func (B BinaryTree) Delete(root *bNode, data interface{}) *bNode {
	panic("implement me")
}

// 利用一个队列
// root先放入
// 出一个节点，再把它的左右节点放进去
// 循环直到队列为空
func (B BinaryTree) LevelOrder(root *bNode) []interface{} {
	if root == nil {
		return nil
	}
	var res []interface{}
	var queue []*bNode
	// root先入队
	queue = append(queue, root)
	for len(queue) > 0 {
		// queue出队一个
		// 把节点的左右子节点再入队即可
		n := queue[0]
		// 干掉前一个
		queue = queue[1:]
		res = append(res, n.data)
		if n.left != nil {
			queue = append(queue, n.left)
		}
		if n.right != nil {
			queue = append(queue, n.right)
		}
	}
	return res
}

func (B BinaryTree) PreOrder(root *bNode, res *[]interface{}) {
	if root == nil {
		return
	}
	*res = append(*res, root.data)
	B.PreOrder(root.left, res)
	B.PreOrder(root.right, res)
	return
}

// 用栈来实现递归: 递去=入栈 归来=出栈
// 前序遍历:节点入栈前访问
// 步骤:
// 1. 从根节点,沿左边节点,依次入栈(递去),直到左节点为空。(可以准备归来了)
// 2. 出栈 ,pop出来的有右节点,把右节点当成子树,回到步骤1;如果没有右节点,回到步骤2
// 3. 直至栈为空
func (B BinaryTree) PreOrder1(root *bNode, res *[]interface{}) {
	if root == nil || res == nil {
		return
	}
	ls := LinkStack{}
	stack, _ := ls.Init(0)
	// 步骤1
	cur := root
	for cur != nil {
		// 节点入栈前访问,下同
		*res = append(*res, cur.data)
		_ = ls.Push(stack, cur)
		cur = cur.left
	}
	for !ls.IsEmpty(stack) {
		n := ls.Pop(stack)
		// 步骤2
		if n.(*bNode).right != nil {
			cur := n.(*bNode).right
			for cur != nil {
				*res = append(*res, cur.data)
				_ = ls.Push(stack, cur)
				cur = cur.left
			}
		}
	}
}

func (B BinaryTree) InOrder(root *bNode, res *[]interface{}) {
	if root == nil {
		return
	}
	B.InOrder(root.left, res)
	*res = append(*res, root.data)
	B.InOrder(root.right, res)
}

// 步骤同前序遍历
// 不同点: 出栈的时候访问数据
func (B BinaryTree) InOrder1(root *bNode, res *[]interface{}) {
	if root == nil || res == nil {
		return
	}
	ls := LinkStack{}
	stack, _ := ls.Init(0)
	// 步骤1
	cur := root
	for cur != nil {
		_ = ls.Push(stack, cur)
		cur = cur.left
	}
	for !ls.IsEmpty(stack) {
		n := ls.Pop(stack)
		// 出栈的时候访问
		*res = append(*res, n.(*bNode).data)
		// 步骤3
		if n.(*bNode).right != nil {
			cur := n.(*bNode).right
			for cur != nil {
				_ = ls.Push(stack, cur)
				cur = cur.left
			}
		}
	}
}

func (B BinaryTree) PostOrder(root *bNode, res *[]interface{}) {
	if root == nil {
		return
	}
	B.PostOrder(root.left, res)
	B.PostOrder(root.right, res)
	*res = append(*res, root.data)
}

func (B BinaryTree) TreeDepth(root *bNode) int {
	if root == nil {
		return 0
	}
	leftDepth := B.TreeDepth(root.left) + 1
	rightDepth := B.TreeDepth(root.right) + 1
	if leftDepth > rightDepth {
		return leftDepth
	}
	return rightDepth
}

func (B BinaryTree) RestoreBTree(i []interface{}, i2 []interface{}) *bNode {
	panic("implement me")
}
