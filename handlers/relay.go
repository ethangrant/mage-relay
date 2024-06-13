package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ethangrant/mage-relay/socket"
	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Relay(relay socket.Relay, w http.ResponseWriter, r *http.Request) {
	count := len(relay.Clients)
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
	}

	relay.Clients[count + 1] = c

	c.SetCloseHandler(func(code int, text string) error {
		fmt.Println("Client closed connection:", code, text, strconv.Itoa(count + 1))
		return nil
	})

	msg := "connection number: " + strconv.Itoa(count + 1)
	err = c.WriteMessage(websocket.TextMessage, ([]byte)(msg))
		if err != nil {
			log.Println("write:", err)
		}
}