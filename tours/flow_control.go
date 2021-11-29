package main

import (
	"fmt"
	"math"
)

func mySqrt(x float64) string {
	if (x < 0) {
		return mySqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(mySqrt(2), mySqrt(-4))
	if num0 := 9; num0 > 0 {
		fmt.Println(num0)
	}
}

/*
	在go中,只有一个循环关键字 -> for
	for 循环的三个组件那里可以没有括号, 循环体大括号不可省略
	循环的第一个和第三个组件可以省略, 分号也可以省略
	循环的第二个组件也可省略,省略后将无限循环
*/
