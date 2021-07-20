package 回溯算法

import (
	"fmt"
	"testing"
)

func TestNQueen_nqueen(t *testing.T) {
	nq := new(NQueen)
	n := 1
	res := nq.nQueen(n)
	fmt.Printf("n:%v res:%+v\nlen:%v\n", n, res, len(res))
	n = 2
	res = nq.nQueen(n)
	fmt.Printf("n:%v res:%+v\nlen:%v\n", n, res, len(res))
	n = 3
	res = nq.nQueen(n)
	fmt.Printf("n:%v res:%+v\nlen:%v\n", n, res, len(res))
	n = 4
	res = nq.nQueen(n)
	fmt.Printf("n:%v res:%+v\nlen:%v\n", n, res, len(res))
	n = 5
	res = nq.nQueen(n)
	fmt.Printf("n:%v res:%+v\nlen:%v\n", n, res, len(res))
	n = 8
	res = nq.nQueen(n)
	fmt.Printf("n:%v res:%+v\nlen:%v\n", n, res, len(res))
}
