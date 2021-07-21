package main

type A struct {
	Name string
}

func main() {
	a := &A{Name: "jiebin"}
	var b, c *A
	b = a
	c = a
	println(b == c)
}

func insert(arr []string, data string) []string {
	arr = append(arr, data)
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i+1] < arr[i] {
			arr[i+1], arr[i] = arr[i], arr[i+1]
		}
	}
	return arr
}
