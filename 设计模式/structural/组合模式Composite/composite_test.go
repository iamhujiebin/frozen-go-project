package 组合模式Composite

import (
	"fmt"
	"testing"
)

func TestTraverseDir(t *testing.T) {
	dirs := make([]*Directory, 4)
	// 12个文件
	for i := 0; i < 4; i++ {
		dirs[i] = NewDirectory()
		for j := 0; j < 3; j++ {
			dirs[i].AddComponent(NewFile(fmt.Sprintf("这是文件%d", i*3+j))) // i*3，就可以让数字连续了
		}
	}
	for i := 1; i < 4; i++ {
		dirs[0].AddComponent(dirs[i])
	}
	dirs[0].Traverse()
}
