package 基数排序

// 正整数排序
func RadixSort(arr []uint) {
	if len(arr) < 2 {
		return
	}
	high := findHighBit(arr)
	var bucket [10][]uint // 0-9的桶
	for i := 1; i <= high; i++ {
		j := i
		mod := uint(1)
		for j > 1 {
			j--
			mod *= 10
		}
		// arr数据入桶
		for _, v := range arr {
			bucket[(v/mod)%10] = append(bucket[(v/mod)%10], v)
		}
		// arr数据出桶
		k := 0
		for i := 0; i < len(bucket); i++ {
			for _, j := range bucket[i] {
				arr[k] = j
				k++
			}
		}
		bucket = [10][]uint{}
	}
}

// 找最高位置
func findHighBit(arr []uint) int {
	if len(arr) < 1 {
		return 0
	}
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	high := 0
	for max > 0 {
		high++
		max /= 10
	}
	return high
}
