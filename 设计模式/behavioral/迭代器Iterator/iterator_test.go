package 迭代器Iterator

import (
	"testing"
)

func TestNewArrayIterator(t *testing.T) {
	iter := NewArrayIterator([]interface{}{1, 2, 3, 4})
	for iter.HasNext() {
		println(iter.Value().(int))
		iter.Next()
	}
	iter.index = 0
	for i := iter.Index(); iter.HasNext(); iter.Next() {
		i = iter.index
		println(i, iter.Value().(int))
	}
}
