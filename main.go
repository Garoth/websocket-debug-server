package main

import (
	"flag"
	"log"
	"net/http"

	"code.google.com/p/go.net/websocket"
	"github.com/Garoth/go-signalhandlers"
)

const (
	HTTP_WEBSOCKET = "/websocket"
)

var (
	ADDR = flag.String("port", ":9217", "listening port")
)

func main() {
	log.SetFlags(log.Ltime)
	flag.Parse()

	go signalhandlers.Interrupt()
	go signalhandlers.Quit()

	http.Handle(HTTP_WEBSOCKET, websocket.Handler(HandleWebSocket))

	log.Println("Starting Websocket Debugging Server")
	if err := http.ListenAndServe(*ADDR, nil); err != nil {
		log.Fatalln("Can't start server:", err)
	}
}

func HandleWebSocket(ws *websocket.Conn) {
	var message string

	defer func() {
		log.Println("Closing connection with", ws.RemoteAddr())
		ws.Close()
	}()

	for {
		if websocket.Message.Receive(ws, &message) != nil {
			break
		} else {
			log.Println(message)
		}
	}
}
