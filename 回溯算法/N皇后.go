package 回溯算法

// N皇后
// N宫格,要求每一行都要放一个皇后,但是之间不能有攻击(皇后是左右上下/斜线都能攻击的)
// 返回的Res格式
// [
//   [[Q...],[..Q.],[...Q]] //可行棋盘1
//   [[Q...],[..Q.],[...Q]] //可行棋盘2
//   [[Q...],[..Q.],[...Q]] //可行棋盘3
//   [[Q...],[..Q.],[...Q]] //可行棋盘4
// ] 的三维字符串数组,Q 放置的皇后位置 "."不放置皇后位置
type NQueen struct {
	Res   [][][]string // 三维数组
	Track [][]string   // 二维数组
	N     int          // 棋盘大小
}

func (n *NQueen) nQueen(N int) [][][]string {
	n.Res = make([][][]string, 0)
	n.Track = make([][]string, 0)
	n.N = N
	n.backtrack()
	return n.Res
}

func (n *NQueen) backtrack() {
	if len(n.Track) == n.N {
		path := make([][]string, 0)
		for k := range n.Track {
			path = append(path, n.Track[k])
		}
		n.Res = append(n.Res, path)
		return
	}
	for col := 0; col < n.N; col++ {
		pos := make([]string, n.N)
		for j := range pos {
			pos[j] = "."
		}
		pos[col] = "Q"
		if n.queenAttack(col) {
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
// 判断逻辑:
//  0.依赖Track数组,长度为0直接返回false
// 	1.Track数组中,当前列的其他行有没有Q(列攻击)
//	2.Track每一行中,col+i|col-i位置有没有Q(斜攻击),k--,i++
// @param col:尝试放置Q的列
// @return true:被攻击 false:不被攻击
func (n *NQueen) queenAttack(col int) bool {
	// 情况0
	if len(n.Track) <= 0 {
		return false
	}
	// 情况1
	for _, row := range n.Track {
		if row[col] == "Q" {
			return true
		}
	}
	// 情况2
	i := 1
	for k := len(n.Track) - 1; k >= 0; k-- {
		if col-i >= 0 && n.Track[k][col-i] == "Q" ||
			col+i < n.N && n.Track[k][col+i] == "Q" {
			return true
		}
		i++
	}
	return false
}
