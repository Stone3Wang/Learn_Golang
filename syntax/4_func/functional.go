package main

//func main() {
//	Func4()
//	Func6Invoke()
//	Func7()
//	ClosureInvoke()
//}

// 函数式编程：方法作为变量
func Func4() {
	myFunc3 := Func3 //方法赋值给某个变量，myFunc3本质是变量，该变量刚好是一个方法
	str, err := myFunc3(1, 2)
	println(str, err)
	_, _ = myFunc3(3, 4) //不要返回值
	//发起方法调用
	_, _ = Func3(5, 6)
	Func5()
}

// 函数式编程：局部方法
func Func5() {
	//作用域仅在本方法内
	//匿名方法
	fn := func(name string) string {
		return "hello, " + name
	}
	str := fn("kitty")
	println(str)
}

// 函数式编程：方法作为返回值
// func:关键字；Func6:方法名；()：方法参数；func(name string)string：返回值是个方法
func Func6() func(name string) string {
	return func(name string) string {
		return "Say good bye, " + name + " !"
	} //返回一个匿名方
}

// 调用上述函数
func Func6Invoke() {
	fn := Func6()      //返回值是一个方法
	str := fn("Trump") //返回的是string
	println(str)       //打印返回值string
}

// 函数式编程：匿名方法发起调用
func Func7() {
	fn := func(name string) string {
		return "See u again " + name + " !"
	}("Bidden") //fn是变量类型
	println(fn)
}
