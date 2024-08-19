package main

import "io"

// 泛型实现求和
func Sum[T Number](vals []T) T {
	var res T
	for _, val := range vals {
		res += val
	}
	return res
}

// 求最大值、最小值
func Max[T Number](vals []T) T {
	t := vals[0]
	for i := 1; i < len(vals); i++ {
		if t < vals[i] {
			t = vals[i]
		}
	}
	return t
}

func Min[T Number](vals []T) T {
	t := vals[0]
	for i := 1; i < len(vals); i++ {
		if t > vals[i] {
			t = vals[i]
		}
	}
	return t
}

// 过滤和查找
func Find[T any](vals []T, filter func(t T) bool) T {
	for _, v := range vals {
		if filter(v) {
			return v
		}
	}
	var t T
	return t
}

// 在切片中插入元素
func Insert[T any](idx int, val T, vals []T) []T {
	if idx < 0 || idx > len(vals) {
		panic("idx 不合法")
	}
	vals = append(vals, val) //在末尾插入，先扩容
	//从后往前
	for i := len(vals) - 1; i > idx; i-- {
		if i-1 >= 0 {
			vals[i] = vals[i-1] //向后移动元素
		}
	}
	vals[idx] = val
	return vals
}

type Interger int
type Number interface {
	~int | uint | int32 //衍生类型
}

func UseSum() {
	res := Sum[int]([]int{123, 123})
	println(res)
	resV1 := Sum[Interger]([]Interger{123, 123})
	println(resV1)
}

func Closable[T io.Closer]() {
	var t T
	t.Close()
}
