package 希尔排序

func ShellSort(arr []int) {
	n := len(arr)
	for step := n / 2; step > 0; step = step / 2 {
		groupSort(arr, 1)
	}
}

func groupSort(arr []int, step int) {
	n := len(arr)
	for i := step; i < n; i += step {
		for j := i - step; j >= 0; j -= step {
			if arr[j+step] < arr[j] {
				arr[j+step], arr[j] = arr[j], arr[j+step]
			}
		}
	}
}