package main

import (
	"fangkm.demo/myclient"
	"fmt"
)

func main() {
	//client := myclient.NewClient()
	//
	//user, err := client.User.CreateUser("fangkiaming")
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//fmt.Printf("%+v", user)
	err :=myclient.NewMyError("错误...")
	fmt.Println(err)
}
