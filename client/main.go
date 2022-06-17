package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	dl := websocket.Dialer{}
	conn, _, err := dl.Dial("ws://127.0.0.1:8888", nil)
	if err != nil {
		log.Println(err)
		return
	}
	go send(conn)
	for {
		m, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println(m, string(p))
	}
}
func send(conn *websocket.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		l, _, _ := reader.ReadLine()
		conn.WriteMessage(websocket.TextMessage, []byte(l))
	}
}
