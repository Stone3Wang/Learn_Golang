package main

import "fmt"

func main() {
	s1, s2, err := DeleteAt([]int{1, 2, 3, 4, 5}, 10)
	fmt.Printf("新切片:%v, 被删除的元素:%v, 错误:%v\n", s1, s2, err)
}
