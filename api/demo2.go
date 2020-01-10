package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//listen, err := net.Listen("tcp", ":5678")
	//if err != nil {
	//	log.Fatalf("Connection failed, error: %v", err)
	//}
	var data = map[string]string{
		"name": "fangkaiming",
		"id":   "ddd",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, res *http.Request) {
		dataJson, _ := json.Marshal(data)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("content-type", "application/json")
		_, _ = w.Write(dataJson)
	})

	fmt.Println("server start...")
	err := http.ListenAndServe(":23456", nil)
	log.Fatal(err)
}
