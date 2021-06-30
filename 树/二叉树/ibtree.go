package 二叉树

// 节点
type BNode struct {
	Data  interface{}
	Left  *BNode
	Right *BNode
}

// 树头
type Root *BNode

// Binary Tree :二叉树
type IBinaryTree interface {
	InitTree(interface{}) Root                      // 初始化树
	LevelOrder(Root) []interface{}                  // 层序遍历
	PreOrder(Root) []interface{}                    // 前序遍历:中左右
	InOrder(Root) []interface{}                     // 中序遍历:左中右
	PostOrder(Root) []interface{}                   // 后序遍历:左右中
	TreeDepth(Root) int                             // 树深度
	RestoreBTree([]interface{}, []interface{}) Root // 还原二叉树
}
