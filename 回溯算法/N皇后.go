package 回溯算法

// N皇后
// N宫格,要求每一行都要放一个皇后,但是之间不能有攻击(皇后是左右上下/斜线都能攻击的)
// 返回的Res格式
// [
//   [[Q...],[..Q.],[...Q]] //可行棋盘1
//   [[Q...],[..Q.],[...Q]] //可行棋盘2
//   [[Q...],[..Q.],[...Q]] //可行棋盘3
//   [[Q...],[..Q.],[...Q]] //可行棋盘4
// ] 的三维整型数组,Q 放置的皇后位置 "."不放置皇后位置
type NQueen struct {
	Res   [][][]string // 三维数组
	Track [][]string   // 二维数组
	N     int          // 棋盘大小
}

func (n *NQueen) nQueen(N int) [][][]string {
	n.Res = make([][][]string, 0)
	n.Track = make([][]string, 0)

	n.N = N
	return n.Res
}

func (n *NQueen) backtrack() {
	if len(n.Track) == n.N {
		path := make([][]string, 0)
		copy(path, n.Track)
		n.Res = append(n.Res, path)
		return
	}
	for i := 0; i < n.N; i++ {
		canPos, pos := n.checkQueenAttack()
		if !canPos {
			continue
		}
		// 做选择
		n.Track = append(n.Track, pos)
		// 决策树
		n.backtrack()
		// 撤销选择
		n.Track = n.Track[0 : len(n.Track)-1]
	}
}

// 检查皇后是否被攻击
// @return canPos:能否放置 true:能 false:不能
// @return pos: 能放置
func (n *NQueen) checkQueenAttack() (canPos bool, pos []string) {
	pos = make([]string, n.N)
	for i := range pos {
		pos[i] = "."
	}
	for i := 0; i < n.N; i++ {
		
	}
	return false, pos
}
