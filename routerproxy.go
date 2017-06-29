package main

import "fmt"
import "sync"
import "time"
import "net/http"
import "io"
import "encoding/json"

import "golang.org/x/crypto/ssh"
import "golang.org/x/net/websocket"

import "github.com/jonstout/routerproxy/database"
import "github.com/jonstout/routerproxy/device"

func testSSH() {
	config := &ssh.ClientConfig{
		User: "",
		Auth: []ssh.AuthMethod{
			ssh.Password(""),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	commands := []*device.Command{
		&device.Command{
			Text: "show version\n",
			Type: "cisco",
			Addr: "156.56.6.220:22",
		},
		&device.Command{
			Text: "show version\n",
			Type: "cisco",
			Addr: "156.56.6.221:22",
		},
		&device.Command{
			Text: "show version\n",
			Type: "juniper",
			Addr: "156.56.6.115:22",
		},
		&device.Command{
			Text: "show version\n",
			Type: "juniper",
			Addr: "156.56.6.116:22",
		},
	}

	var wg sync.WaitGroup
	gt := time.Now()

	for _, cmd := range commands {

		wg.Add(1)
		go func(cmd *device.Command) {
			defer wg.Done()
			t := time.Now()

			if cmd.Type == "cisco" {
				c := device.NewCisco(config)
				_, _ = c.Execute(cmd)

				//fmt.Println(string(bytes))
				fmt.Println(time.Since(t))
			} else {
				c := device.NewJuniper(config)
				_, _ = c.Execute(cmd)

				//fmt.Println(string(bytes))
				fmt.Println(time.Since(t))
			}
		}(cmd)
	}

	wg.Wait()
	fmt.Println(time.Since(gt))
}

func echo(ws *websocket.Conn) {
	io.Copy(ws, ws)
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

    http.Handle("/ws", websocket.Handler(echo))
    http.ListenAndServe(":8000", nil)
}
