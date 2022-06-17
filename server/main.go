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
		log.Println("upgrade err:", err)
		return
	}
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		fmt.Println(string(p))
	}
	defer conn.Close()
	log.Println("服务关闭")
}
func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe("localhost:8888", nil)
}
