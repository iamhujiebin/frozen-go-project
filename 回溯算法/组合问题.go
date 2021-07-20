package 回溯算法

// 组合问题
// 例子: [1,2,3] 选2个--> [[1 2] [1 3] [2 3]]
// 解题思路: 回溯算法
// 画棵树
/*
		  O
	  1/  2\  3\
      O    O    O
     2/3| 3|
    O  O  O
*/
type Combination struct {
	Res   [][]int
	Track []int // 这里存储arr的index (区别于全排列,存储arr的元素值;主要为了判断j > i)
	N     int   // 组合中选的个数
}

func (c *Combination) Combine(arr []int, n int) [][]int {
	c.Res = make([][]int, 0)
	c.N = n
	c.Track = make([]int, 0)
	c.backtrack(arr)
	return c.Res
}

func (c *Combination) backtrack(arr []int) {
	if len(c.Track) == c.N {
		track := make([]int, 0)
		for _, index := range c.Track {
			track = append(track, arr[index])
		}
		c.Res = append(c.Res, track)
		return
	}
	for index := 0; index < len(arr); index++ {
		// 决策中的数,需要比前面的index要大,不然会出现重复(例如: [1,2] [2,1])
		if index <= getMax(c.Track) {
			continue
		}
		// 做选择
		c.Track = append(c.Track, index)
		// 决策树
		c.backtrack(arr)
		// 撤销选择(removeLast)
		c.Track = c.Track[0 : len(c.Track)-1]
	}
}

func getMax(arr []int) int {
	max := -1
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}
