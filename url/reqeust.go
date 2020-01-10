package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":9001",
	}
	http.HandleFunc("/users", user)
	_ = server.ListenAndServe()
}

type User struct {
	Name    string   `json:"name"`
	Friends []string `json:"friends"`
}

func user(w http.ResponseWriter, r *http.Request) {
	var user = &User{
		Name:    "fangkaiming",
		Friends: []string{"lili", "xiaxia"},
	}

	jsonData, _ := json.Marshal(user)
	w.Header().Set("content-type","application/json")
	w.Write(jsonData)
}
