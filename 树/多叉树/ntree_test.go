package 多叉树

import "testing"

var nt NTree
var ntree *nNode

func init() {
	ntree = nt.InitTree()
}

func TestNTree_TreeDepth(t *testing.T) {
	t.Log(nt.TreeDepth(ntree))
}

func TestNTree_PreOrder(t *testing.T) {
	t.Log(nt.PreOrder(ntree))
}

func TestNTree_PostOrder(t *testing.T) {
	t.Log(nt.PostOrder(ntree))
}

func TestNTree_LevelOrder(t *testing.T) {
	t.Log(nt.LevelOrder(ntree))
}

func TestNTree_Find(t *testing.T) {
	t.Log(nt.Find(ntree, 1))
	t.Log(nt.Find(ntree, 11))
	t.Log(nt.Find(ntree, 111))
}

func TestFindLowParent(t *testing.T) {
	t.Log(FindLowParent(ntree, 4, 6))
	t.Log(FindLowParent(ntree, 4, 7))
	t.Log(FindLowParent(ntree, 2, 11))
}

func TestNTree_FindWithParent(t *testing.T) {
	t.Log(nt.FindWithParent(ntree, 1))
	t.Log(nt.FindWithParent(ntree, 5))
	t.Log(nt.FindWithParent(ntree, 4))
	t.Log(nt.FindWithParent(ntree, 11))
}
