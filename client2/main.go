package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8888", nil)
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			l, _, _ := reader.ReadLine()
			conn.WriteMessage(websocket.TextMessage, l)
		}
	}()
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
}
