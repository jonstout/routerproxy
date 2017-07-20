package main

import "fmt"
import "encoding/json"
import "io/ioutil"
import "strings"
import "time"

import "github.com/jonstout/routerproxy/database"
import "github.com/jonstout/routerproxy/device"
import "github.com/gorilla/websocket"

import "golang.org/x/crypto/ssh"


type Credential struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

func HandleCommands(conn *websocket.Conn) {
	file, err := ioutil.ReadFile("./conf/config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		var req device.CommandRequest
		err := conn.ReadJSON(&req)
		if err != nil {
			fmt.Println(err)
			break
		}

		var cred Credential
		err = json.Unmarshal(file, &cred)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, id := range req.Devices {
			config := &ssh.ClientConfig{
				User: cred.Username,
				Auth: []ssh.AuthMethod{
					ssh.Password(cred.Password),
				},
				HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			}

			go func(config *ssh.ClientConfig, deviceID int) {
				fmt.Println("Running", req.Command.Selection)

				d := database.GetDevice(deviceID)
				fmt.Println(d)
				address := strings.Join([]string{d.IPAddress, ":22"}, "")

				var res device.CommandResponse
				res.Command = req.Command
				res.DeviceID = d.ID

				t := time.Now()

				if d.Type == "juniper" {
					c := device.NewJuniper(config)
					bytes, err := c.Execute(address, &req.Command)
					if err != nil {
						res.Output = err.Error()
					} else {
						res.Output = string(bytes)
					}
				} else {
					c := device.NewCisco(config)
					bytes, err := c.Execute(address, &req.Command)
					if err != nil {
						res.Output = err.Error()
					} else {
						res.Output = string(bytes)
					}
				}

				fmt.Println(time.Since(t))
				conn.WriteJSON(&res)
			}(config, id)
		}
	}
}
