package main

import "fmt"
//import "sync"
//import "time"
import "net/http"
//import "io"
import "encoding/json"

//import "golang.org/x/crypto/ssh"
//import "golang.org/x/net/websocket"
import "github.com/gorilla/websocket"

import "github.com/jonstout/routerproxy/database"
import "github.com/jonstout/routerproxy/device"

// func testSSH() {
// 	config := &ssh.ClientConfig{
// 		User: "",
// 		Auth: []ssh.AuthMethod{
// 			ssh.Password(""),
// 		},
// 		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
// 	}

// 	commands := []*device.Command{
// 	}

// 	var wg sync.WaitGroup
// 	gt := time.Now()

// 	for _, cmd := range commands {

// 		wg.Add(1)
// 		go func(cmd *device.Command) {
// 			defer wg.Done()
// 			t := time.Now()

// 			if cmd.Type == "cisco" {
// 				c := device.NewCisco(config)
// 				_, _ = c.Execute(cmd)

// 				//fmt.Println(string(bytes))
// 				fmt.Println(time.Since(t))
// 			} else {
// 				c := device.NewJuniper(config)
// 				_, _ = c.Execute(cmd)

// 				//fmt.Println(string(bytes))
// 				fmt.Println(time.Since(t))
// 			}
// 		}(cmd)
// 	}

// 	wg.Wait()
// 	fmt.Println(time.Since(gt))
// }

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

	go func() {
		for {
			var req device.CommandRequest
			err := conn.ReadJSON(&req)
			if err != nil {
				fmt.Println(err)
				break
			}

			fmt.Println(req)

			for _, deviceID := range req.Devices {
				var res device.CommandResponse
				res.Command = req.Command
				res.DeviceID = deviceID
				res.Output = "a\nb\nc\n"

				conn.WriteJSON(&res)
			}
		}
	}()
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, `{"device": "switch.test.com"}`)
}

func devices(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	devices := database.Devices()
	encoder.Encode(devices)
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/device", devices)
    http.HandleFunc("/ws", wsHandler)

    http.ListenAndServe(":8000", nil)
}
