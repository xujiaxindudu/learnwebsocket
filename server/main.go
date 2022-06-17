package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var UP = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handle(w http.ResponseWriter, r *http.Request) {
	conn, err := UP.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		m, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println(m, string(p))
	}
	defer conn.Close()
	log.Println("服务关闭")
}
func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe("localhost:8888", nil)
}
