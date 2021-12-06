package main

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p) //42
	fmt.Println(p)
	*p = 21
	fmt.Println(i) //21
	p = &j
	*p = *p / 37
	fmt.Println(j) //73
}
