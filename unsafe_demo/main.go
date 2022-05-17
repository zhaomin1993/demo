package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var x struct {
		a bool
		b int16
		c []int
	}

	type y struct {
		ya bool
		yb int16
		yc []int
	}

	/**
	  unsafe.Offsetof 函数的参数必须是一个字段 x.f, 然后返回 f 字段相对于 x 起始地址的偏移量, 包括可能的空洞.
	*/

	/**
	  uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	  指针的运算
	*/
	// 和 pb := &x.b 等价
	x.a = true
	x.b = 10
	y_ := (*y)(unsafe.Pointer(&x))
	fmt.Println(*y_)
	pa := (*bool)(unsafe.Pointer(&x))
	*pa = true
	fmt.Println(x.a)
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b) // "42"
	pc := (*[]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.c)))
	*pc = []int{1, 2, 3}
	fmt.Println(x.c)
}

// string2bytes 字符串转byte切片
func string2bytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

//bytes2string byte切片转字符串
func bytes2string(b []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}
	return *(*string)(unsafe.Pointer(&sh))
}

// String byte切片转字符串
func String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Str2Bytes 字符串转byte切片
func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
