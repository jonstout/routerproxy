package main

import "fmt"
import "encoding/json"
import "net/http"

import "github.com/jonstout/routerproxy/database"
import "github.com/gorilla/mux"
import "github.com/gorilla/websocket"


func devices(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	devices := database.GetDevices()
	encoder.Encode(devices)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println(err)
        return
    }

	go HandleCommands(conn)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/device", devices)
	r.HandleFunc("/ws", wsHandler)

	http.ListenAndServe(":8000", r)
}
