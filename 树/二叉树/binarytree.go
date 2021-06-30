package 二叉树

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
func (B BinaryTree) InitTree(datas []interface{}) Root {
	//if len(datas) <= 0 {
	//	return nil
	//}
	n1, n2, n3, n4, n5, n6, n7, n8, n9, n10 :=
		&BNode{Data: 1}, &BNode{Data: 2}, &BNode{Data: 3}, &BNode{Data: 4}, &BNode{Data: 5},
		&BNode{Data: 6}, &BNode{Data: 7}, &BNode{Data: 8}, &BNode{Data: 9}, &BNode{Data: 0}
	n1.Left, n1.Right = n2, n3
	n2.Left, n2.Right = n4, n5
	n3.Left = n6
	n4.Left, n4.Right = n7, n8
	n5.Left, n5.Right = n9, n10
	return n1
}

// 利用一个队列
// root先放入
// 出一个节点，再把它的左右节点放进去
// 循环直到队列为空
func (B BinaryTree) LevelOrder(root Root) []interface{} {
	if root == nil {
		return nil
	}
	var res []interface{}
	var queue []*BNode
	// root先入队
	queue = append(queue, root)
	for len(queue) > 0 {
		// queue出队一个
		// 把节点的左右子节点再入队即可
		node := queue[0]
		// 干掉前一个
		queue = queue[1:]
		res = append(res, node.Data)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return res
}

func (B BinaryTree) PreOrder(root Root) []interface{} {
	if root == nil {
		return nil
	}
	var res []interface{}

}

func (B BinaryTree) InOrder(root Root) []interface{} {
	panic("implement me")
}

func (B BinaryTree) PostOrder(root Root) []interface{} {
	panic("implement me")
}

func (B BinaryTree) TreeDepth(root Root) int {
	if root == nil {
		return 0
	}
	leftDepth := B.TreeDepth(root.Left) + 1
	rightDepth := B.TreeDepth(root.Right) + 1
	if leftDepth > rightDepth {
		return leftDepth
	}
	return rightDepth
}

func (B BinaryTree) RestoreBTree(i []interface{}, i2 []interface{}) Root {
	panic("implement me")
}
