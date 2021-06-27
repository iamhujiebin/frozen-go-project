package 迭代器Iterator

// 迭代器模式
// 自己实现遍历而已
type Iterator interface {
	Index() int
	Value() interface{}
	HasNext() bool
	Next()
}

type ArrayIterator struct {
	index  int
	values []interface{}
}

func (a *ArrayIterator) Index() int {
	return a.index
}

func (a *ArrayIterator) Value() interface{} {
	return a.values[a.index]
}

func (a *ArrayIterator) HasNext() bool {
	if a.index < len(a.values) {
		return true
	}
	return false
}

func (a *ArrayIterator) Next() {
	if a.HasNext() {
		a.index++
	}
}

func NewArrayIterator(arr []interface{}) *ArrayIterator {
	return &ArrayIterator{
		index:  0,
		values: arr,
	}
}
