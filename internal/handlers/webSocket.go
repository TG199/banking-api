package handlers

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/TG199/banking-api/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	clients  = make(map[*websocket.Conn]bool)
	broadcast = make(chan models.Account)
	mutex = &sync.Mutex{}
)

func HandleConnections(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}

	defer conn.Close()

	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	for {
		var msg models.Account
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Error reading JSON:", err)
			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()
			break
		}
	}
}

func BroadcastBalanceUpdate(account models.Account) {
    broadcast <- account
}
func HandleMessages() {
	for {
		account := <-broadcast
		mutex.Lock()
		for client := range clients {
			err := client.WriteJSON(account)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}
