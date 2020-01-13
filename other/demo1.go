package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var str = "方开明"
	fmt.Println("run method: ", len([]rune(str)))
	fmt.Println("utf8:", utf8.RuneCountInString(str))

	var i interface{} = 99
	j := i.(int)

	fmt.Printf("%T=>%d\n", j, j)

	if i, ok := i.(int); ok {
		fmt.Printf("%T => %d\n", i, i)
	}

	fmt.Println(strings.Repeat("-", 20))
	testFunc(10, "string", 1.2)
}

func testFunc(item ...interface{}) {
	for i, v := range item {
		switch v.(type) {
		case string:
			fmt.Printf("param #%d is a string\n", i)
		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case float64:
			fmt.Printf("param #%d is a float64\n", i)
		case int, int8, int16, int32, int64:
			fmt.Printf("param #%d is a int\n", i)
		case nil:
			fmt.Printf("param #%d is nil\n", i)
		default:
			fmt.Printf("param #%d is unkown\n", i)
		}
	}
}
