package main

import "fmt"

func ChangeUser() {
	fmt.Printf("------------u1---------------\n")
	u1 := User{Name: "John", Age: 18, Nickname: "Johnny"}
	fmt.Printf("%+v\n", u1)
	fmt.Printf("u1 address:%p\n", &u1)
	u1.ChangeName("Jerry")
	u1.ChangeAge(20)
	fmt.Printf("%+v\n", u1)
	fmt.Printf("u1 address:%p\n", &u1)

	fmt.Printf("------------u2---------------\n")

	u2 := User{Name: "Tom", Age: 25, Nickname: "Tommy"}
	fmt.Printf("%+v\n", u2)
	fmt.Printf("u2 address:%p\n", &u2)
	u1.ChangeName("Kitty")
	u1.ChangeAge(35)
	fmt.Printf("%+v\n", u2)
	fmt.Printf("u2 address:%p\n", &u2)
}

type User struct {
	Age      int
	Name     string
	Nickname string
}

// 结构体接收器：值传递，修改的是副本
func (u User) ChangeName(name string) {
	u.Nickname = name
	fmt.Printf("ChangeName:user address:%p\n", &u)
}

// 指针结构体接收器
func (u *User) ChangeAge(age int) {
	u.Age = age
}
