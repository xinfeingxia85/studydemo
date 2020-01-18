package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
)

func main() {
	//byte := make([]byte, 10)
	//_, _ = rand.Read(byte)
	//fmt.Println(string(byte))

	data := make([]byte, 1024)
	rand.Read(data)
	hash := md5.New()
	hash.Write(data)
	//hash.Sum(nil)
	fmt.Println(hex.EncodeToString(hash.Sum(nil)))
}
