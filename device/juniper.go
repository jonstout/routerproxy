package device

import "bufio"
import "log"
import "regexp"
import "time"

import "golang.org/x/crypto/ssh"

type Juniper struct {
	config *ssh.ClientConfig
}

func NewJuniper(config *ssh.ClientConfig) *Juniper {
	return &Juniper{config}
}

func (c *Juniper) Execute(addr string, cmd *Command) ([]byte, error) {
	client, err := ssh.Dial("tcp", addr, c.config)
	if err != nil {
		log.Fatal(err)
	}

	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	input, _ := session.StdinPipe()
	output, _ := session.StdoutPipe()

	if err := session.Shell(); err != nil {
		log.Fatal(err)
	}

	newLine := make(chan []byte)

	go func() {
		cli, _ := regexp.Compile(`.*@.*>.*`)
		screen, _ := regexp.Compile(`Screen length`)

		reader := bufio.NewReader(output)

		for {
			bytes, err := reader.ReadBytes([]byte("\n")[0])
			if err != nil {
				break
			}

			ok := cli.Match(bytes)
			if ok {
				continue
			}

			ok = screen.Match(bytes)
			if ok {
				continue
			}

			newLine <- bytes
		}
	}()

	input.Write([]byte("set cli screen-length 0\n"))
	input.Write(append([]byte(cmd.Selection), "\n"...))

	bytes := make([]byte, 0, 1024)
	for {
		select {
		case l := <-newLine:
			bytes = append(bytes, l...)
		case <-time.After(time.Millisecond * 500):
			session.Close()
			return bytes, err
		}
	}
}
