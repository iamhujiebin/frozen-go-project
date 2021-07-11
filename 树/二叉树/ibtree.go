package 二叉树

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

// 节点
type bNode struct {
	data  interface{}
	left  *bNode
	right *bNode

	height int // avl 平衡二叉搜索树用到

	isBlack bool   // 红黑树用到,是否黑色节点
	parent  *bNode // 红黑树父节点
}

func (n *bNode) ToInt() int {
	i, ok := n.data.(int)
	if ok {
		return i
	}
	return INT_MIN
}

// Binary Tree :二叉树
type IBinaryTree interface {
	InitTree(...interface{}) *bNode                   // 初始化树
	Find(*bNode, interface{}) (*bNode, *bNode)        // 查找节点
	Insert(*bNode, interface{}) *bNode                // 插入节点
	Delete(*bNode, interface{}) *bNode                // 删除节点
	LevelOrder(*bNode) []interface{}                  // 层序遍历
	PreOrder(*bNode, *[]interface{})                  // 前序遍历:中左右
	InOrder(*bNode, *[]interface{})                   // 中序遍历:左中右
	PostOrder(*bNode, *[]interface{})                 // 后序遍历:左右中
	TreeDepth(*bNode) int                             // 树深度
	RestoreBTree([]interface{}, []interface{}) *bNode // 还原二叉树
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
