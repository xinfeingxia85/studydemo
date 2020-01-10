package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Name string    `json:"name"`
	Sex  string    `json:"sex"`
	Date time.Time `json:"date",omitempty`
}

func main() {
	user := User{
		Name: "fangkm",
		Sex:  "男",
		Date: time.Now(),
	}

	//jsonData, _ := json.Marshal(user)
	jsonData, _ := json.MarshalIndent(user, "", "  ")
	fmt.Printf("%s\n", jsonData)
}
