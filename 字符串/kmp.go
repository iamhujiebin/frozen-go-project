package 字符串

import "strings"

// 返回子串索引
func IndexGo(s, substr string) int {
	return strings.Index(s, substr)
}

// 暴力算法brute force解决
func IndexBF(s, substr string) int {
	m, n := len(s), len(substr)
	if m < n {
		return -1
	}
	i, i2, j := 0, 0, 0
	for i < m && j < n {
		if s[i2] == substr[j] {
			i2++
			j++
		} else { // 不匹配,都回溯 i=>i+1 j=>0
			i++
			i2 = i
			j = 0
		}
	}
	// 匹配到了
	if j == n {
		return i
	}
	return -1
}

// KMP看毛片算法
// 求next数组
// 匹配失败时: s不需要回溯,substr按照next数组回溯
// 注意bf是返回i,KMP是返回i-n:因为i不回溯,找到之后是匹配串的最后一个位置
func IndexKMP(s, substr string) int {
	//next := nextArr(substr)
	next := nextValueArr(substr) // 用nextValue数组优化
	m, n := len(s), len(substr)
	i, j := 0, 0
	for i < m && j < n {
		if s[i] == substr[j] {
			i++
			j++
		} else {
			j = next[j]
			// 只要next数组对应值不是-1
			// i不动,j动
			// 否则 i++ && j=0
			if j == -1 {
				i++
				j = 0
			}
		}
	}
	// 匹配到了
	if j == n {
		return i - n // 返回第一个位置
	}
	return -1
}

// important: 求next数组
// next数组就是KMP模型,匹配失败后,子串回溯的位置,公式如下
// 公式：求匹配失败位置"前面"子子串的最长"前后缀"匹配长度再+1
// 注意,由于数组的下标从0开始,所以上面的+1 再-1 = 最长公共前后缀的长度
func nextArr(substr string) []int {
	n := len(substr)
	if n <= 0 {
		return nil
	}
	next := make([]int, n)
	if n == 1 {
		next[0] = -1
		return next
	}
	if n == 2 {
		next[0], next[1] = -1, 0
		return next
	}
	next[0], next[1] = -1, 0 // 不必计算，每个都一样的
	for i := 2; i < n; i++ {
		match := 0 // 最大公共前缀
		// 比较 str[0:j] str[i-j:i]
		ss := substr[0:i] // 取出失败前的字串
		for j := 1; j < i; j++ {
			h := ss[0:j]
			t := ss[i-j : i]
			if h == t {
				if j > match {
					match = j
				}
			}
		}
		next[i] = match
	}
	return next
}

// next数组的优化
// 思路:假如next数组指向的位置的字符,跟匹配失败位置的字符一样,则失败位置的next数组的值,可以直接用回溯位置next数组的值。
// 简称:用"值中值"
func nextValueArr(substr string) []int {
	n := len(substr)
	if n <= 0 {
		return nil
	}
	next := make([]int, n)
	if n == 1 {
		next[0] = -1
		return next
	}
	if n == 2 {
		next[0], next[1] = -1, 0
		return next
	}
	next[0], next[1] = -1, 0 // 不必计算，每个都一样的
	for i := 2; i < n; i++ {
		match := 0        // 最大公共前缀
		ss := substr[0:i] // 取出失败前的字符串
		for j := 1; j < i; j++ {
			// 比较 ss[0:j] ss[i-j:i]
			h := ss[0:j]
			t := ss[i-j : i]
			if h == t {
				if j > match {
					match = j
				}
			}
		}
		// 如果匹配失败的字符,跟准备回溯位置的字符一致
		// 那么:next数组值直接沿用
		// 否则:next数组值不变
		if substr[i] == substr[match] {
			next[i] = next[match] // 这里其实不用考虑前面再前面的位置了,因为按顺序来,就会只存在最前面的位置
		} else {
			next[i] = match
		}
	}
	return next
}
