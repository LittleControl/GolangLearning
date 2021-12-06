package main

import "fmt"

type Vertex struct {
	// X int
	// Y int
	X, Y int
}

func main() {
	// fmt.Println("Start")
	// fmt.Println(Vertex{1, 2})
	// v := Vertex{1, 2}
	// v.X = 3
	// v.Y = 4
	// fmt.Println(v)
	// fmt.Println(v.X)
	// p := &v
	// p.X = 5
	// fmt.Println(v.X)
	// fmt.Println(p.X)
	v1 := Vertex{1, 2}
	v2 := Vertex{X: 3}
	v3 := Vertex{}
	p := &Vertex{4, 5}
	fmt.Println(v1, v2, v3, p)
	fmt.Println(p.X)
}
