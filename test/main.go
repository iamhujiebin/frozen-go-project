package main

func main() {
	const INT_MAX = int(^uint(0) >> 1)
	const INT_MIN = ^INT_MAX
	println(INT_MAX, INT_MIN)
}
