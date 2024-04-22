package main

import (
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn  *websocket.Conn
	Mutex sync.Mutex
}

var (
	clients    = sync.Map{}                
	broadcast  = make(chan []byte, 100)    
	workerPool = make(chan chan []byte, 10) 
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true 
	},
}

func main() {
	setupWorkers(10) 

	go handleMessages()
	go connectToBinance()

	http.HandleFunc("/ws", handleConnections)
	log.Println("HTTP server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func setupWorkers(count int) {
	for i := 0; i < count; i++ {
		worker := make(chan []byte, 100)
		go messageWorker(worker)
		workerPool <- worker 
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	client := &Client{Conn: ws}
	clients.Store(ws, client)

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			log.Printf("error: %v", err)
			clients.Delete(ws)
			break
		}
	}
}

func handleMessages() {
	for msg := range broadcast {
		worker := <-workerPool
		worker <- msg
		workerPool <- worker 
	}
}

func messageWorker(messages chan []byte) {
	for msg := range messages {
		clients.Range(func(key, value interface{}) bool {
			client := value.(*Client)
			client.Mutex.Lock()
			err := client.Conn.WriteMessage(websocket.TextMessage, msg)
			client.Mutex.Unlock()
			if err != nil {
				log.Printf("error: %v", err)
				client.Conn.Close()
				clients.Delete(key)
				return true
			}
			return true
		})
	}
}

func connectToBinance() {
	var addr = "stream.binance.com:9443"
	u := url.URL{Scheme: "wss", Host: addr, Path: "/ws/btcusdt@trade"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	log.Println("connected to Binance")

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		broadcast <- message
	}
}
