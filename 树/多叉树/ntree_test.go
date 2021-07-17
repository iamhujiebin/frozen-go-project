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
