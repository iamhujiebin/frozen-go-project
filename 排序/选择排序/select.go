package 选择排序

func SelectSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[minIndex], arr[i] = arr[i], arr[minIndex]
		}
	}
}

func SelectSort2(arr []int, minIndex int) {
	n := len(arr)
	if minIndex >= n {
		return
	}
	tmpMinIndex := minIndex
	for i := minIndex; i < n; i++ {
		if arr[i] < arr[tmpMinIndex] {
			tmpMinIndex = i
		}
	}
	if tmpMinIndex != minIndex {
		arr[tmpMinIndex], arr[minIndex] = arr[minIndex], arr[tmpMinIndex]
	}
	SelectSort2(arr, minIndex+1)
}

// 同时把最小最大找出来
// 递归方式
func SelectSort3(arr []int, minIndex, maxIndex int) {
	if minIndex >= maxIndex {
		return
	}
	tmin, tmax := minIndex, maxIndex
	for i := minIndex; i < maxIndex; i++ {
		if arr[i] < arr[tmin] {
			tmin = i
		}
		if arr[i] > arr[tmax] {
			tmax = i
		}
	}
	if tmin != minIndex {
		arr[tmin], arr[minIndex] = arr[minIndex], arr[tmin]
	}
	if tmax != maxIndex {
		arr[tmax], arr[maxIndex] = arr[maxIndex], arr[tmax]
	}
	SelectSort3(arr, minIndex+1, maxIndex-1)
}

func SelectSort4(arr []int) {
	if len(arr) < 2 {
		return
	}
	minIndex, maxIndex := 0, len(arr)-1
	for minIndex < maxIndex {
		tmin, tmax := minIndex, maxIndex
		for i := minIndex; i < maxIndex; i++ {
			if arr[i] < arr[tmin] {
				tmin = i
			}
			if arr[i] > arr[tmax] {
				tmax = i
			}
		}
		if tmin != minIndex {
			arr[tmin], arr[minIndex] = arr[minIndex], arr[tmin]
		}
		if tmax != maxIndex {
			arr[tmax], arr[maxIndex] = arr[maxIndex], arr[tmax]
		}
		minIndex++
		maxIndex--
	}
}
