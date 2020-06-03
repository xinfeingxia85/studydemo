package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't excute this on a window machine")
	} else {
		execute()
	}
}

func execute() {
	out, err := exec.Command("ls").Output()

	//if there is an error with our execution
	//handle it here
	if err != nil {
		fmt.Printf("%s", err)
	}

	// as the out variable defined above is of type []byte we need to convert
	//this to a string  or we will see garbage printed out in the console
	//this is how we convert it to a string
	fmt.Println("Command sucessfully executed")
	output := string(out[:])
	fmt.Println(output)

	//let's try the pwd command here
	out, err = exec.Command("pwd").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Command sucessfully executed")
	output = string(out[:])
	fmt.Println(output)
}
