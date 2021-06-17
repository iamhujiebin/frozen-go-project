package 冒泡排序

func BubbleSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}
	for i := n - 1; i >= 0; i-- {
		shift := false
		for j := 0; j < i; j++ {
			if arr[j+1] < arr[j] {
				shift = true
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
		if !shift {
			return
		}
	}
}

func BubbleSort2(arr []int, n int) {
	if n < 2 {
		return
	}
	shift := false
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			shift = true
			arr[i+1], arr[i] = arr[i], arr[i+1]
		}
	}
	if !shift {
		return
	}
	BubbleSort2(arr, n-1)
}
