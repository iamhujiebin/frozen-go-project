package intx

// 是否包含int
func IntIndexOf(list []int, target int) int {
	index := -1
	for i, v := range list {
		if v == target {
			index = i
		}
	}
	return index
}
