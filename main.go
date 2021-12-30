package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func paginaInicial(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Pagina Inicial")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Cliente Conectado")

	err = ws.WriteMessage(1, []byte("Olá Cliente!"))
	if err != nil {
		log.Println(err)
	}
	reader(ws)

}

func setandoRotas() {
	http.HandleFunc("/", paginaInicial)
	http.HandleFunc("/ws", wsEndpoint)

}

func main() {
	fmt.Println("Bem vindo ao WebSocket")
	setandoRotas()

	log.Fatal(http.ListenAndServe(":8080", nil))
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
