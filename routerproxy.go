package main

import "fmt"
import "sync"
import "time"

import "golang.org/x/crypto/ssh"

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

func main() {
	testSSH()
}
