package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	var users Users
	jsonFile, err := os.Open("usrs.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	err = json.NewDecoder(jsonFile).Decode(&users)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Printf("%+v", users)
	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
		fmt.Println("---------------------------")
	}
}
