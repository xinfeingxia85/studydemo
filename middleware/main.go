package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.LstdFlags|log.Lshortfile)
	server := http.Server{
		Addr: ":9003",
	}

	m := MiddleWare{}

	http.HandleFunc("/test", m.panic(m.log(hadleWeb)))

	server.ListenAndServe()
}

func hadleWeb(w http.ResponseWriter, r *http.Request) {
	panic("xxxx")
	fmt.Fprint(w, "测试")
}
