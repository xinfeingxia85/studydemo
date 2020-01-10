package main

import (
	"log"
	"net/http"
	"time"
)

type MiddleWare struct {
}

func (m MiddleWare) log(handlerFunc http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		handlerFunc(w, r)
		t2 := time.Now()
		diff := t2.Sub(t1)
		log.Printf("[%s] %s: %v", r.Method, r.URL.String(), diff)
	}

	return fn
}

func (m MiddleWare) panic(handlerFunc http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Print("[panic] recoverd from :", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		handlerFunc(w, r)
	}

	return  fn
}
