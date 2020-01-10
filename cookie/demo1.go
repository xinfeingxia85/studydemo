package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":9002",
	}

	http.HandleFunc("/setcookie", setCookie)
	http.HandleFunc("/getcookie", getCookie)
	_ = server.ListenAndServe()
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	c1, _ := r.Cookie("zhulixia")

	_, _ = fmt.Fprintf(w, "%s=%s\n", c1.Name,c1.Value)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:  "fangkaiming",
		Value: "112",
	}

	c2 := http.Cookie{
		Name:  "zhulixia",
		Value: "12233",
	}

	w.Header().Set("Set-Cookie", c1.String())
	http.SetCookie(w, &c2)
	_, _ = fmt.Fprintf(w, "%s\n", c1.String())
}
