package main

import "fmt"

func Slice() {
	s1 := []int{1, 2, 3, 4} //直接初始化4个元素的切片
	fmt.Printf("s1:%v,len:%d,cap:%d \n", s1, len(s1), cap(s1))

	s2 := make([]int, 3, 4) //make函数创建切片，指定长度和容量，长度为3，容量为4
	fmt.Printf("s2:%v,len:%d,cap:%d \n", s2, len(s2), cap(s2))

	s2 = append(s2, 7) //使用append函数添加元素到切片中,没有扩容
	fmt.Printf("append第一次后：s2:%v,len:%d,cap:%d \n", s2, len(s2), cap(s2))

	s2 = append(s2, 8) //扩容了，容量变为8
	fmt.Printf("appennd第二次后：s2:%v,len:%d,cap:%d \n", s2, len(s2), cap(s2))

	s3 := make([]int, 4) //长度和容量都为4
	fmt.Printf("s3:%v,len:%d,cap:%d \n", s3, len(s3), cap(s3))
	//在初始化时要预估容量,遇事不决用make函数初始化切片
	//推荐写法
	//s4 := make([]int, 0, 4) //长度为0，容量为4
}

// 子切片
func SubSlice() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := s1[3:6] //子切片，左闭右开
	fmt.Printf("s2:%v,len:%d,cap:%d \n", s2, len(s2), cap(s2))

	s3 := s1[:6] //子切片，左闭右开
	s4 := s1[3:] //子切片，左闭右开
	fmt.Printf("s3:%v,len:%d,cap:%d \n", s3, len(s3), cap(s3))
	fmt.Printf("s4:%v,len:%d,cap:%d \n", s4, len(s4), cap(s4))

}

// 共享数组
func ShareSlice() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[2:] //子切片

	fmt.Printf("s1:%v,len:%d,cap:%d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%v,len:%d,cap:%d \n", s2, len(s2), cap(s2))

	println("取子切片改元素值后:\n")
	s2[0] = 99
	fmt.Printf("s1:%v,len:%d,cap:%d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%v,len:%d,cap:%d \n", s2, len(s2), cap(s2))

	println("扩容第一次:\n")
	s2 = append(s2, 200)
	fmt.Printf("s1:%v,len:%d,cap:%d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%v,len:%d,cap:%d \n", s2, len(s2), cap(s2))

	println("扩容第二次:\n")
	s2 = append(s2, 300)
	fmt.Printf("s1:%v,len:%d,cap:%d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%v,len:%d,cap:%d \n", s2, len(s2), cap(s2))

}
