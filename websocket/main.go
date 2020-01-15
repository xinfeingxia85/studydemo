package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	fmt.Println("ddd")
	setRoutes()
	log.Fatal(http.ListenAndServe(":9005", nil))
}

func setRoutes() {
	//http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	content := `
    <!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <title>Go WebSocket Tutorial</title>
  </head>
  <body>
    <h2>Hello World</h2>

    <script>
        let socket = new WebSocket("ws://127.0.0.1:9005/ws");
        console.log("Attempting Connection...");

        socket.onopen = () => {
            console.log("Successfully Connected");
            socket.send("Hi From the Client!")
        };
        
        socket.onclose = event => {
            console.log("Socket Closed Connection: ", event);
            socket.send("Client Closed!")
        };

        socket.onerror = error => {
            console.log("Socket Error: ", error);
        };

    </script>
  </body>
</html>
    `
	_, _ = w.Write([]byte(content))
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	//whether an incoming request from a different domain is allowd to connect,
	// and if it isn't they'll be hit with a CORS error
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	//upgrade this connection to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	//helpfu log statement to show connection
	log.Println("Client connected")
	err = ws.WriteMessage(1,[]byte("Hi client"))
	if err !=nil{
		log.Println(err)
	}

	reader(ws)
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
