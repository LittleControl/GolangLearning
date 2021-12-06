package main

import (
	"fmt"
	// "math"
	"runtime"
	"time"
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
	// fmt.Println(mySqrt(2))
	// fmt.Println(math.Sqrt(2))
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s. \n", os)
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon.")
	default:
		fmt.Println("Good Evening")
	}

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

/*
go 在switch里如果遇到符合条件的case会默认break
case的值不能是一个constant的变量,值不是必须是int
*/
