package main

func Loop1() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	//也可以这样写,在循环体中写递归式
	for i := 0; i < 10; {
		println(i)
		i++
	}
}

func Loop2() {
	i := 0
	for i < 10 {
		println(i)
		i++
	}
}

// 无限循环，cpu 100%
func Loop3() {
	for {
		println("hello")
	}
}

func ForArray() {
	println("遍历数组")
	arr := [3]string{"A", "B", "C"}
	for idx, val := range arr {
		println(idx, val)
	}
	//用下标
	for idx := range arr {
		println(idx, arr[idx])
	}
}

func ForSlice() {
	println("遍历切片")
	arr := []string{"A", "B", "C"}
	for idx, val := range arr {
		println(idx, val)
	}
	//忽略下标只用值
	for _, val := range arr {
		println(val)
	}

}

func ForMap() {
	println("遍历map")
	m := map[string]string{
		"1": "A",
		"2": "B",
	}
	for k, v := range m {
		println(k, v)
	}

	println("遍历key,忽略value")
	for k := range m {
		println(k, m[k])
	}
}

func Loopbug() {
	users := []User{
		{Name: "Alice"},
		{Name: "Bob"},
	}
	m := make(map[string]*User)
	//不要对迭代参数取地址
	for _, user := range users {
		m[user.Name] = &user
	}
	for name, user := range m {
		println(name, user.Name)
	}
}

type User struct {
	Name string
}

func ForBreak() {
	for i := 0; i < 20; i++ {
		if i >= 10 {
			break
		}
		println("For里面", i)
	}

}

func ForContinue() {
	for i := 0; i < 10; i++ {
		println("Continue 前", i)
		if i%2 == 0 {
			continue
		}
		println("Continue 后", i)
	}
}
