package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &DemoHandle{})

	server := &http.Server{
		Addr:    ":5678",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}

type DemoHandle struct{}

func (*DemoHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "apllication/json")
	//w.Write([]byte("测试了"))
	fmt.Fprintln(w, "哈哈哈")
}
