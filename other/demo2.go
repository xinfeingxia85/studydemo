package main

import (
	"fmt"
)

type Test struct {
}

func main() {
	a := new(Test)
	fmt.Printf("%P\n", a)

	b := new([]int)
	fmt.Println(b)

	m := make([]int, 10, 50)
	fmt.Println(m)
}
