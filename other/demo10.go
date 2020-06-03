package main

import (
	"fmt"
	"regexp"
)

func main() {
	var department []string
	dn := "cn=fangkaiming,ou=架构技术组,ou=产研中心,ou=C端业务群,ou=三节课,ou=users,dc=sanjieke,dc=com"
	reg := regexp.MustCompile(`ou=([^,]*)`)
	results := reg.FindAllStringSubmatch(dn, -1)
	num := len(results)
	department = make([]string, num-1)
	num = num-1
	for key, sub := range results {
		if key == (len(results) - 1) {
			continue
		}
		num--
		department[num] = sub[1]
	}

	//fmt.Println(department)
	fmt.Printf("%#v",department)
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
