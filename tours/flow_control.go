package main

import (
	"fmt"
	"math"
)

// z -= (z*z - x) / (2*z)
func mySqrt(x float64) float64 {
	var res float64 = 1
	for count := 0; count < 10; count++ {
		res -= (res * res - x) / (2 * res)
		fmt.Println(res)
	}
	return res
}

func main() {
	fmt.Println(mySqrt(2))
	fmt.Println(math.Sqrt(2))
}

/*
	在go中,只有一个循环关键字 -> for
	for 循环的三个组件那里可以没有括号, 循环体大括号不可省略
	循环的第一个和第三个组件可以省略, 分号也可以省略
	循环的第二个组件也可省略,省略后将无限循环
*/

/*
if 和 else中条件中定义的变量,只能在自己的块范围内引用
*/
