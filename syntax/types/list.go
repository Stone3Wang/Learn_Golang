package main

// 接口定义
// 接口是一组行为的抽象
type List interface {
	Add(index int, value any)
	Append(value any)
	Delete(index int)
}

// 结构体定义
type LinkedList struct {
	head node
}

func (l *LinkedList) Append(value any) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Delete(index int) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Add(index int, value any) {
	//实现这个方法
}

// 结构体自引用
type node struct {
	next *node
}

func UseListV1() {
	l := &LinkedList{}
	l.Add(0, 1)
	l.Add(1, "123")
	l.Add(2, nil)
}
