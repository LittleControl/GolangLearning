package main

import (
	"fmt"
	// "math/cmplx"
	// "math"
)

func add(x , y int) int {
	return x + y
}

func swap(x, y string)(string, string) {
	return y, x
}

func split(sum int)(x, y int) {
	x = sum * 4 - 3
	y = sum * 4 - 4
	return
}

const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x * 10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

// go's basic types
/*
	bool
	string
	int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr
	byte //alias for uint8
	rune //alias for int32 //represents a Unicode code point
	float32 float64
	complex64 complex128

*/
/*
	go 里的零值(或者说初始值)
	指得是当你声明一个变量时,未给定初始值时的默认值
	数字类型默认值为 0
	布尔类型默认值为 false
	字符串类型默认值为 ""
*/

/*
	go 常量可以用关键字 const 来声明
	常量可以是char, string, boolean, or 数字 格式的值
	常量不可以用 := 来声明
*/
