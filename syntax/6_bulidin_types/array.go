package main

import "fmt"

func Array() {
	//直接初始化一个三元素数组
	a1 := [3]int{1, 2, 3}
	fmt.Printf("a1:%v, len:%d, cap:%d\n", a1, len(a1), cap(a1))

	//少了部分元素，自动补零
	a2 := [3]int{4, 5}
	fmt.Printf("a2:%v, len:%d, cap:%d\n", a2, len(a2), cap(a2))

	//没有初始化，默认值是0
	a3 := [3]int{}
	fmt.Printf("a3:%v, len:%d, cap:%d\n", a3, len(a3), cap(a3))
	//数组不支持append操作，只能通过切片来进行扩容
	//a3 = append(a3, 6)

	//下标越界会panic
}

func arr1(index int) {
	a := [3]int{1, 2, 3}
	fmt.Printf("a1[1]:%d\n", a[index])
}
