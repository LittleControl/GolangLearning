package main

import (
	"fmt"
)

func main() {
	sum := 1
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum)
	for ; sum > 2; {
		sum -= 1
	}
	fmt.Println(sum)
	for sum < 5 {
		sum += 1
	}
	fmt.Println("The Wa Lu Du")
	fmt.Println(sum)
}

/*
	在go中,只有一个循环关键字 -> for
	for 循环的三个组件那里没有括号, 循环体大括号不可省略
	循环的第一个和第三个组件可以省略, 分号也可以省略
	循环的第二个组件也可省略,省略后将无限循环
*/
