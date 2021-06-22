package 计数排序

func CountSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	min, max := findMinMax(arr)
	bucket := make([]int, max-min+1)
	for _, v := range arr {
		bucket[v-min]++
	}
	k := 0
	for i := 0; i < len(bucket); i++ {
		for j := 1; j <= bucket[i]; j++ {
			arr[k] = min + i
			k++
		}
	}
}

func findMinMax(arr []int) (min, max int) {
	if len(arr) <= 0 {
		return
	}
	min, max = arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return
}
