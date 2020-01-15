package main

import (
	"fangkm.demo/chat/pkg/websocket"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	fmt.Println("Distributed Chat APP v0.01")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":9006", nil))
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Simple server")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serverWs(pool, w, r)
	})
}

func serverWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		_, _ = fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}
