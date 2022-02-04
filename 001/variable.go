package main

// rune 类似于其他语言的char类型

import (
	"fmt"
	"math/cmplx"
)

func getInitValues() {
	var name string
	var age int
	fmt.Println(name, age)
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	fmt.Println(cmplx.Pow(1i, 2))
}

func enums() {
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)
	fmt.Println(cpp, java, python, golang)
	const (
		javascript = iota
		_
		ruby
	)
	fmt.Println(javascript, ruby)
}

// go的类型转换是强制的

func main() {
	fmt.Println("Hello World")
	getInitValues()
	euler()
	enums()
}
