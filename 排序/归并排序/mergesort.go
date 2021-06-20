package 归并排序

func Merge(a1, a2 []int) (m []int) {
	if len(a1) <= 0 {
		m = a2
		return
	}
	if len(a2) <= 0 {
		m = a1
		return
	}
	l1, l2 := 0, 0
	n1, n2 := len(a1), len(a2)
	for l1 < n1 && l2 < n2 {
		if a1[l1] < a2[l2] {
			m = append(m, a1[l1])
			l1++
		} else {
			m = append(m, a2[l2])
			l2++
		}
	}
	if l1 < n1 {
		m = append(m, a1[l1:]...)
	}
	if l2 < n2 {
		m = append(m, a2[l2:]...)
	}
	return
}

// 递归
func MergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	MergeSort(arr, l, mid)
	MergeSort(arr, mid+1, r)

	// 归来时解决问题
	// [l:mid] [mid+1:r] 两个数组进行合并排序
	// golang左闭右开，并且r-l要大于0,"右开!"
	tmp := Merge(arr[l:mid+1], arr[mid+1:r+1])
	if len(tmp) > 0 {
		for i := 0; i <= r-l; i++ {
			arr[l+i] = tmp[i]
		}
	}
}