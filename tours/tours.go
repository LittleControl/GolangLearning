package main

import (
	"fmt"
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

var c, python, java bool = true, false, false

func main() {
	fmt.Println("Nothing can be done")
	// fmt.Println(time.Now())
	// fmt.Println("My favorite number is", rand.Intn(10))
	// fmt.Println(math.Sqrt(4), math.Sqrt(2))
	// fmt.Println(math.Pi)
	// fmt.Println(add(1, 2))
	// a, b := swap("a", "b")
	// fmt.Println(a, b)
	fmt.Println(split(1))
	var i  = 6
	name := "nothing"
	fmt.Println(i, c, python, java, name)
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
