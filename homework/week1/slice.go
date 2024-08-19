package main

import (
	"errors"
	"fmt"
)

var ErrIndexOutOfRange = errors.New("index out of range")

// DeleteAt 删除指定位置的元素
// 如果下标不是合法的下标，返回 ErrIndexOutOfRange错误

func DeleteAt[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	if index < 0 || index >= length {
		var zero T
		return nil, zero, newErrIndexOutOfRange(length, index)
	}
	res := src[index]
	// 从index位置开始，后面的元素都向前移动一位
	for i := index; i+1 < length; i++ {
		src[i] = src[i+1]
	}
	// 去掉最后一个重复元素
	src = src[:length-1]
	src = Shrink(src)
	return src, res, nil
}

func newErrIndexOutOfRange(length int, index int) error {
	return fmt.Errorf("ekit: %w,下标超出范围, 长度为 %d, 下标为 %d", ErrIndexOutOfRange, length, index)
}

func calCapacity(c, l int) (int, bool) {
	if c <= 64 {
		return c, false
	}
	if c > 2048 && (c/l >= 2) {
		factor := 0.625
		return int(float32(c) * float32(factor)), true
	}
	if c <= 2048 && (c/l >= 4) {
		return c / 2, true
	}
	return c, false
}

// Shrink 缩容
func Shrink[T any](src []T) []T {
	c, l := cap(src), len(src)
	n, changed := calCapacity(c, l)
	if !changed {
		return src
	}
	s := make([]T, 0, n)
	s = append(s, src...)
	return s
}
