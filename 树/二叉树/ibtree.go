package 二叉树

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

// 节点
type BNode struct {
	Data  interface{}
	Left  *BNode
	Right *BNode
}

func (n *BNode) ToInt() int {
	i, ok := n.Data.(int)
	if ok {
		return i
	}
	return INT_MIN
}

// Binary Tree :二叉树
type IBinaryTree interface {
	InitTree(...interface{}) *BNode                   // 初始化树
	Insert(*BNode, *BNode)                            // 插入节点
	Delete(*BNode, *BNode)                            // 删除节点
	LevelOrder(*BNode) []interface{}                  // 层序遍历
	PreOrder(*BNode, *[]interface{})                  // 前序遍历:中左右
	InOrder(*BNode, *[]interface{})                   // 中序遍历:左中右
	PostOrder(*BNode, *[]interface{})                 // 后序遍历:左右中
	TreeDepth(*BNode) int                             // 树深度
	RestoreBTree([]interface{}, []interface{}) *BNode // 还原二叉树
}
