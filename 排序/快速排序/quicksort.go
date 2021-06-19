package 快速排序

// 找中心点pivot,左小右大
// 挖坑填数
// 递归分治, 左右边界记住golang的 左闭右开
func QuickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	if r-l <= 10 {
		insertSort(arr, l, r)
		return
	}
	min, max := l, r
	pivot := arr[l]
	i := l // 坑
	moving := 1
	for l < r {
		if moving%2 == 1 {
			if arr[r] < pivot {
				arr[i], arr[r] = arr[r], arr[i]
				i = r
				moving++
				continue
			}
			r--
		}
		if moving%2 == 0 {
			if arr[l] > pivot {
				arr[i], arr[l] = arr[l], arr[i]
				i = l
				moving++
				continue
			}
			l++
		}
	}
	arr[l] = pivot
	QuickSort(arr, min, l)
	QuickSort(arr, l+1, max)
}

// 优化点
func insertSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	for i := l + 1; i <= r; i++ {
		for j := i - 1; j >= l; j-- {
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}
