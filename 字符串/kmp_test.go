package 字符串

import "testing"

var str = "xxxxxababcxxxxx"
var substr = "ababc"

func TestIndexGo(t *testing.T) {
	t.Log(IndexGo(str, substr))
}

func TestIndexBF(t *testing.T) {
	t.Log(IndexBF(str, substr))
	t.Log(IndexBF("a", substr))
	t.Log(IndexBF("ababac", substr))
}

func TestNextArr(t *testing.T) {
	t.Log(nextArr(substr))
	t.Log(nextValueArr(substr))
	t.Log(nextArr("hello"))
	t.Log(nextValueArr("hello"))
	t.Log(nextArr("aaaab"))
	t.Log(nextValueArr("aaaab"))
}

func TestIndexKMP(t *testing.T) {
	t.Log(IndexGo(str, substr))
	t.Log(IndexKMP(str, substr))
	t.Log(IndexGo("abdabc", "bc"))
	t.Log(IndexKMP("abdabc", "bc"))
	t.Log(IndexGo("ababc", "abc"))
	t.Log(IndexKMP("ababc", "abc"))
	t.Log(IndexGo("abcabc", "d"))
	t.Log(IndexKMP("abcabc", "d"))
	t.Log(IndexGo("aaaaab", "ab"))
	t.Log(IndexKMP("aaaaab", "ab"))
}
