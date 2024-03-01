package main

//基本类型与string
import (
	"math"
	"strconv"
	"unicode/utf8"
)

func main() {
	//常规运算
	var a int = 100
	var b int = 200
	println(a + b)
	println(a - b)
	println(a * b)
	println(a / b)
	println(float64(a) / float64(b))

	println("float64(a) / float64(b) =", strconv.FormatFloat(float64(a)/float64(b), 'f', 3, 64))
	a++
	println(a)

	println(math.Atan(1))

	println("a+b=" + strconv.Itoa(a+b))

	//Extremum()
	//String()
	//Byte()
	Bool()
}

// 极值
func Extremum() {
	println(math.MinInt64)
	println("float64 最小正数", math.SmallestNonzeroFloat64)
	println("float32 最小正数", math.SmallestNonzeroFloat32)
}

// string类型
func String() {
	println("Hello " +
		"go")
	println('h')
	//he said :"hello world"
	println("he said :\"hello world\"")

	//len用法
	println(len("hello world"))                    //输出11
	println(len("hello world你好"))                  //输出17
	print(utf8.RuneCountInString("hello world你好")) //输出13，一个中文汉字代表一个字节
}

// byte类型
func Byte() {
	var a byte = 'a'
	println(a) //输出a的ASCII值97

	var str string = "this is string"
	var bs []byte = []byte(str) //类型转换
	println(bs)
}

// bool类型
func Bool() {
	var a bool = true
	var b bool = false
	println(a && b)
	println(a || b)
	println(!a)
	//组合取反
	//!(a && b) = !a || !b
	//!(a || b) = !a && !b
	println(!(a && b))
	println(!(a || b))
}
