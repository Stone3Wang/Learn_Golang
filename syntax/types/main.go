package main

import "fmt"

func main() {
	u1 := &User{} //取地址
	println(u1)
	u1 = new(User) //取地址
	println(u1)

	u2 := User{} //取值
	fmt.Printf("%+v\n", u2)
	//修改字段值
	u2.Name = "Jerry"
	u2.Age = 25
	fmt.Printf("%+v\n", u2)

	var u3 User
	fmt.Printf("%+v\n", u3) //声明但没初始化

	var u4 *User
	println(u4) //nil

	u5 := User{Name: "Tom", Age: 30}
	fmt.Printf("%+v\n", u5)
	ChangeUser()
	components()
}

// 结构体和指针
func UseList() {
	l1 := LinkedList{}
	l1Ptr := &l1
	var l2 LinkedList = *l1Ptr
	fmt.Printf("%+v\n", l2)

	//这是nil
	var l3Ptr *LinkedList
	println(l3Ptr)
}
