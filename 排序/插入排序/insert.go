package 插入排序

func InsertSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}
	for i := 0; i < n; i++ {
		index := i
		value := arr[i] // 找扑克牌位置
		for j := i; j >= 0; j-- {
			if arr[j] > arr[i] {
				index = j
			}
		}
		if index != i {
			copy(arr[index+1:i+1], arr[index:i+1]) // golang语法,左闭右开
			arr[index] = value
		}
	}
}
