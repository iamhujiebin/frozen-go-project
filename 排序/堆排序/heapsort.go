package 堆排序

// 下沉元素
// 最大堆,最大的在顶
func heapify(arr []int, dad int, end int) {
	son := dad*2 + 1
	for son <= end {
		if son+1 <= end && arr[son+1] > arr[son] {
			son++
		}
		if arr[son] > arr[dad] {
			arr[son], arr[dad] = arr[dad], arr[son]
		}
		dad = son
		son = dad*2 + 1
	}
}

// 递归下沉
func heapify1(arr []int, dad int, end int) {
	son := dad*2 + 1
	if son > end {
		return
	}
	if son+1 <= end && arr[son+1] > arr[son] {
		son++
	}
	if arr[son] > arr[dad] {
		arr[dad], arr[son] = arr[son], arr[dad]
	}
	heapify1(arr, son, end)
}

func HeapSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	// 初始化堆
	// 从最后一个父节点开始下沉(heapify)
	n := len(arr) - 1
	for i := (n - 1) / 2; i >= 0; i-- {
		heapify(arr, i, n)
	}

	// 堆顶最大,和最后一个元素换位置,再堆化
	for i := n; i >= 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapify1(arr, 0, i-1) //再建堆, end=i-1就是 最后一个元素的位置再减去1
	}
}
