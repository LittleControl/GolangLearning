package main

import (
	"fmt"
)

func switchTest(score int) string {
	res := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong Score: %d", score))
	case score < 90:
		res = "B"
	case score <= 100:
		res = "B"
	}
	return res
}

func main() {
	const filename = "variable.go"
	// contents, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%s\n", contents)
	// }
	// if contents, err := ioutil.ReadFile(filename); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%s\n", contents)
	// }
	fmt.Printf("%s, %s, %s", switchTest(100), switchTest(70), switchTest(-1))
}
