package main

// defer执行顺序： 后进先出
func Defer() {
	defer func() {
		println("第一个defer")
	}()
	defer func() {
		println("第二个defer")
	}()
}

// defer与闭包
// 作为闭包引入
func DeferClosure() {
	k := 0
	defer func() {
		println(k)
	}()
	k = 2
}

// 作为参数传入
func DeferClosureV1() {
	m := 0
	defer func(m int) {
		println(m)
	}(m)
	//defer func(val int) {
	//	println(val)
	//}(m)
	//将变量m的值传入到参数val中
	m = 3
	//m的值与func(m int)中的传入的参数无关
	// defer func()中的参数要与}后()中的参数同时存在
}

// defer修改返回值
func DeferReturn() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferReturnV1() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return a
}

// 指针形式,成功更改返回值：返回值是指针，更改的是指针指向的值
func DeferReturnV2() *MyStruct {
	res := &MyStruct{
		name: "Tommy",
	}
	defer func() {
		res.name = "Jerry"
	}()
	return res
}

type MyStruct struct {
	name string
}

//看看变更多少
