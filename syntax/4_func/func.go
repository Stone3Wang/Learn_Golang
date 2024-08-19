package main

func main() {
	Invoke()
	//Recursive()
	//Defer()
	//DeferClosure()
	//DeferClosureV1()
	////fmt.Println("DeferReturn", DeferReturn())
	////fmt.Println("DeferReturnV1", DeferReturnV1())
	////fmt.Println("DeferReturnV2", DeferReturnV2().name)
	//DeferClosureLoopV1()
	//DeferClosureLoopV2()
	//DeferClosureLoopV3()
}

// 接收返回值
func Invoke() {
	println(Func0("wl"))
	str1, err := Func1(1, 2, 3, "wl") //:=表示err是新变量
	println(str1, err)
	//忽略返回值
	_, err = Func3(4, 5) //不加:表示err已经定义
	println(err)
	str2, _ := Func3(7, 8) //加:表示str2是新变量需要使用:=
	println(str2)
}

// 方法声明：返回值
// Func0单一返回值
func Func0(name string) string {
	return "hello, " + name
}

// 多个返回值:较多
func Func1(a, b, c int, d string) (string, error) {
	return "", nil
}

// Func2带名字的返回值
func Func2(a int, b int) (str string, err error) {
	str = "hello"
	//带名字的返回值，可以直接return
	return
}

// Func3带名字的返回值
func Func3(a int, b int) (str string, err error) {
	res := "hello"
	//带名字的返回值，但实际没使用
	return res, nil
}

// 方法调用：递归
func Recursive() {
	Recursive()
}
