package main

import "basic_Go/syntax/3_variables/demo"

var (
	a1 int     = 456
	a2 float64 = 3.14
)

func main() {
	var a int = 123 //完整的变量命名，局部变量
	println(a)

	//var (
	//	a1 int     = 456
	//	a2 float64 = 3.14
	//)
	//类型推断:整数默认int，浮点数默认float64类型
	//相同作用域下只可声明一次变量，不同可声明多次
	var a1 = "123"
	println(a1 + "hello") //不要用这种写法

	var b = 789
	println(b)
	var c = 2.14156
	println(c)

	var d uint = 88
	println(d)

	var s = "string"
	println(s)

	var f int
	println(f)

	println(demo.Global)

	i := 100
	println(i)

	//常量声明iota用法
	const (
		Status0 = iota + 2
		Status1
		Status2
		a
		//插入主动赋值中断iota
		Status3 = 10
		Status4
	)
	//每次左移1位
	const (
		status0 = iota<<10 + 1
		status1
		status2
	)

}
