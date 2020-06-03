package main

import (
	"fmt"
	"sort"
)

func main() {
	str :=[]string{"1","7", "8", "10"}
	fmt.Println(sort.SearchStrings(str, "78"))
}
