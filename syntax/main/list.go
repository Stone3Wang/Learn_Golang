package main

type ListV1[T any] interface {
	Add(index int, value T)
	Append(val T) error
	Delete(index int) error
}

// 泛型结构体
type LinkedListV1[T any] struct {
	head *nodeV1[T]
}

func (l *LinkedListV1[T]) Add(index int, val T) {

}

type nodeV1[T any] struct {
	data T
}

func UseList() {
	l := &LinkedListV1[int]{}
	l.Add(1, 123)
	//l.Add(1,."123")
}
