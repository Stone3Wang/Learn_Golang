package main

import "fmt"

type Fish struct {
}

func (f Fish) Swim() {

}
func (f Fakefish) FakeSwim() {

}

type Fakefish Fish

func UseFish() {
	f1 := Fish{}
	f1.Swim()

	f2 := Fakefish{}
	f2.FakeSwim()

	//类型转换
	f3 := Fakefish(f1)

	f4 := Fish(f2)

	fmt.Println(f3, f4)

	//类型别名
	type Yu = Fish
	y := Yu{}
	y.Swim()
}
