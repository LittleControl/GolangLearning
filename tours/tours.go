package main

import (
	"fmt"
	// "math/cmplx"
	"math"
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



func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x * x + y * y))
	var z uint = uint(f)
	var d = uint(f)
	var k = f
	// var s int = f
	fmt.Println(x, y, z, d, k)
	var s = 3
	var t = 3.0
	fmt.Printf("%T, %v\n", s, s)
	fmt.Printf("%T, %v\n", t, t)
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
