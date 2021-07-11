package tree

type Element interface {
	Value() interface{}
	Compare(Element) int //相等返回0，小于返回负数，大于返回正数
}
