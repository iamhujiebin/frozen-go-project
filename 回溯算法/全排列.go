package 回溯算法

import "frozen-go-project/utils/intx"

// 求全排列
// 例子: [1,2,3] --> [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
// 解题思路: 回溯算法
// 画棵树
/*
		  O
	  1/  2\  3\
      O    O    O
     2/3| 1/3| 1/ 2\
    O  O  O   O O   O
    3| 2|  3|  1|  2|  1|
    O   O   O   O   O   O
*/
type Permutation struct {
	Res   [][]int // 结果集
	Track []int   // 中间子过程
}

func (p *Permutation) Permute(arr []int) (res [][]int) {
	p.Res = make([][]int, 0)
	p.Track = make([]int, 0)
	p.backtrack(arr)
	return p.Res
}

func (p *Permutation) backtrack(arr []int) {
	if len(p.Track) == len(arr) {
		track := make([]int, len(p.Track))
		copy(track, p.Track)
		p.Res = append(p.Res, track)
		return
	}
	for _, v := range arr {
		if intx.IntIndexOf(p.Track, v) > -1 {
			continue
		}
		// 做选择
		p.Track = append(p.Track, v)
		// 进入下层决策树
		p.backtrack(arr)
		// 撤销选择(removeLast)
		p.Track = p.Track[0 : len(p.Track)-1]
	}
}
