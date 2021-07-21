package 多叉树

import "frozen-go-project/utils/queuestack"

type nNode struct {
	Data  interface{}
	Nodes []*nNode
}

// n叉树
type INTree interface {
	InitTree(...interface{}) *nNode                      // 初始化树
	Find(*nNode, interface{}) *nNode                     // 查找节点
	FindWithParent(*nNode, interface{}) (*nNode, *nNode) // 查找父子节点
	LevelOrder(*nNode) []interface{}                     // 层序遍历
	PreOrder(*nNode) []interface{}                       // 前序遍历:根子
	PostOrder(*nNode) []interface{}                      // 后序遍历:左右中
	TreeDepth(*nNode) int                                // 树深度
}

type NTree struct {
}

/*
			5
		/   |    \
       1    2     3
      /\    |   / | \
     4 6    7   8 9 10
	/
   11
*/
func (N NTree) InitTree(...interface{}) *nNode {
	n1, n2, n3, n4, n5, n6, n7, n8, n9, n10, n11 := &nNode{Data: 1}, &nNode{Data: 2}, &nNode{Data: 3}, &nNode{Data: 4},
		&nNode{Data: 5}, &nNode{Data: 6}, &nNode{Data: 7}, &nNode{Data: 8}, &nNode{Data: 9}, &nNode{Data: 10}, &nNode{Data: 11}
	n5.Nodes = make([]*nNode, 3)
	n5.Nodes[0], n5.Nodes[1], n5.Nodes[2] = n1, n2, n3
	n1.Nodes = make([]*nNode, 2)
	n1.Nodes[0], n1.Nodes[1] = n4, n6
	n2.Nodes = make([]*nNode, 1)
	n2.Nodes[0] = n7
	n3.Nodes = make([]*nNode, 3)
	n3.Nodes[0], n3.Nodes[1], n3.Nodes[2] = n8, n9, n10
	n4.Nodes = make([]*nNode, 1)
	n4.Nodes[0] = n11
	return n5
}

func (N NTree) Find(root *nNode, data interface{}) *nNode {
	if root == nil {
		return nil
	}
	if root.Data == data {
		return root
	}
	var res *nNode
	for i := range root.Nodes {
		res = N.Find(root.Nodes[i], data)
		if res != nil {
			return res
		}
	}
	return res
}

func (N NTree) FindWithParent(root *nNode, data interface{}) (*nNode, *nNode) {
	panic("todo")
}

func (N NTree) LevelOrder(root *nNode) []interface{} {
	if root == nil {
		return nil
	}
	res := make([]interface{}, 0)
	var lq queuestack.LinkQueue
	queue, _ := lq.Init(0)
	lq.Push(queue, root)
	for !lq.IsEmpty(queue) {
		n := lq.Pop(queue)
		res = append(res, n.(*nNode).Data)
		for i := range n.(*nNode).Nodes {
			lq.Push(queue, n.(*nNode).Nodes[i])
		}
	}
	return res
}

func (N NTree) PreOrder(root *nNode) []interface{} {
	if root == nil {
		return nil
	}
	res := make([]interface{}, 0)
	res = append(res, root.Data)
	for i := range root.Nodes {
		res = append(res, N.PreOrder(root.Nodes[i])...)
	}
	return res
}

func (N NTree) PostOrder(root *nNode) []interface{} {
	if root == nil {
		return nil
	}
	res := make([]interface{}, 0)
	for i := range root.Nodes {
		res = append(res, N.PostOrder(root.Nodes[i])...)
	}
	res = append(res, root.Data)
	return res
}

func (N NTree) TreeDepth(root *nNode) int {
	if root == nil {
		return 0
	}
	var maxDepth int
	for i := range root.Nodes {
		if N.TreeDepth(root.Nodes[i]) > maxDepth {
			maxDepth = N.TreeDepth(root.Nodes[i])
		}
	}
	return 1 + maxDepth
}
