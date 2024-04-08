package main

func main() {
	println(IfNewVar(100, 250))
	println(IfElseIf(15))
}

func IfOnly(age int) string {
	if age >= 18 {
		return "成年了"
	}
	return "未成年"
}

func IfElse(age int) string {
	if age >= 18 {
		return "成年了"
	} else {
		return "未成年"
	}
}

func IfElseIf(age int) string {
	if age >= 18 {
		return "成年了"
	} else if age >= 12 {
		return "少年"
	}
	return "未成年"
}

// 新变量只作用于if-else块中
func IfNewVar(start int, end int) string {
	if distance := end - start; distance > 100 {
		println(distance)
		return "太远了不去"
	} else {
		println(distance)
		return "距离较近"
	}
	//编译错误
	//println(distance)
}
