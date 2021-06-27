package 二叉树

// 节点
type BNode struct {
	Data  interface{}
	Left  *BNode
	Right *BNode
}

// 树头
type Head *BNode

type IBtree interface {
	LevelOrder(Head) []interface{}   // 层序遍历
	PreOrder(Head) []interface{}     // 前序遍历:中左右
	InOrder(Head) []interface{}      // 中序遍历:左中右
	PostOrder(Head) []interface{}    // 后序遍历:左右中
	TreeDepth(Head) int              // 树深度
	RestoreBTree([]interface{}) Head // 还原二叉树
}
