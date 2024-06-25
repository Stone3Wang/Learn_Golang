package main

//结构体实现接口
import (
	"fmt"
	"io"
)

type NameI interface {
	Name() string
}
type Outer struct {
	//Outer实现了NameI接口
	Inner
}
type Outer1 struct {
	*Inner
}

func (i Outer1) Name() string {
	return "outer"
}

type Inner struct {
}

func (i Inner) Name() string {
	return "inner"
}

type Outer2 struct {
	//组合接口
	io.Closer
}

func (i Inner) Hello() {
	fmt.Println("Hello from ", i.Name())
}

func components() {
	var o Outer
	o.Hello()
	//初始化指针
	o1 := Outer1{
		Inner: &Inner{},
	}
	o1.Hello()
}
