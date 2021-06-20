package main

import "fmt"

func main() {
	m := make([]int, 0)
	test(&m)
	fmt.Println(m)
}

func test(m *[]int) {
	*m = append(*m, 1)
	*m = append(*m, 1)
	*m = append(*m, 1)
}
