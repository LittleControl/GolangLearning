package main

import (
	"fmt"
	// "strings"
)

func main() {
	// names := [4]string {
	// 	"little",
	// 	"control",
	// 	"avalon",
	// 	"nothing",
	// }
	// fmt.Println(names)
	// a := names[0:2]
	// b := names[1:3]
	// fmt.Println(a, b)

	// b[0] = "Test"
	// fmt.Println(a, b)
	// fmt.Println(names)
	// q := []int{2, 3, 5, 7, 11, 13}
	// fmt.Println(q)

	// s := []struct {
	// 	name string
	// 	gender bool
	// } {
	// 	{"little", true},
	// 	{"control", false},
	// 	{"nothing", false},
	// }
	// fmt.Println(s)
	// var a [10]int
	// fmt.Println(a[0:10])
	// fmt.Println(a[:10])
	// fmt.Println(a[0:])
	// fmt.Println(a[:])
	// s := []int{2, 3, 5, 7, 11, 13}
	// printSlice(s)
	// s = s[1:3]
	// printSlice(s)
	// var s []int
	// fmt.Println(s, len(s), cap(s))
	// if s == nil {
	// 	fmt.Println("nil!")
	// }
	// a := make([]int)
	// fmt.Println(a)
	// s_of_slice := [][]string {
	// 	[]string{"", "o", ""},
	// 	[]string{"o", "", "o"},
	// 	[]string{"", "o", ""},
	// }
	// s_of_slice[0][0] = "X"
	// s_of_slice[0][2] = "X"
	// s_of_slice[1][1] = "X"
	// s_of_slice[2][0] = "X"
	// s_of_slice[2][2] = "X"
	// for i := 0; i < len(s_of_slice); i++ {
	// 	fmt.Printf("%s\n", strings.Join(s_of_slice[i], " "))
	// }
	var s []int
	fmt.Println(s)
	t := append(s, 1)
	fmt.Println(t)
	t = append(s, 2)
	fmt.Println(t)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}


/*
	slice只是对于底层数组的一个引用,修改相关的内容同时也就后修改相应的底层数组
	对应的其他slice也会同步修改
*/
