package main

import (
	"fmt"
)

func main() {
	var arr1 [3]*string
	arr2 := [...]*string{new(string),new(string),new(string)}
	*arr2[0] = "Red"
	*arr2[1] = "Blue"
	*arr2[2] = "Green"
	arr1 = arr2

	fmt.Println(arr1)

	for k,v := range arr1{
		fmt.Printf("arr[%d] = %s\n", k, *v)
	}

}
