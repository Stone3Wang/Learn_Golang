package main

var str = "Jobs"

// 函数式编程入门：闭包 方法+上下文
func Closure(name string) func() string {
	//闭包
	//name 变量和 str变量是外部Closure闭包外部定义的，属于上下文
	//内部定义的不是上下文
	//方法本身
	return func() string {
		hello := "helo world!"
		return "Hi," + name + " it's the Closure Function:这是闭包！And u " + str + " ? " + hello
	} //方法
}
func ClosureInvoke() {
	c := Closure("Bidden") //方法类型变量
	println(c())
}
