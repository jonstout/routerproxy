package device

import "bufio"
import "log"
import "regexp"
import "time"

import "golang.org/x/crypto/ssh"

type Cisco struct {
	config *ssh.ClientConfig
}

func NewCisco(config *ssh.ClientConfig) *Cisco {
	return &Cisco{config}
}

func (c *Cisco) Execute(cmd *Command) ([]byte, error) {
	client, err := ssh.Dial("tcp", cmd.Addr, c.config)
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
		cli, _ := regexp.Compile(`.*@.*#.*`)

		reader := bufio.NewReader(output)

		for {
			bytes, err := reader.ReadBytes([]byte("\n")[0])
			if err != nil {
				break
			}

			ok := cli.Match(bytes)
			if !ok {
				newLine <- bytes
			}
		}
	}()

	input.Write([]byte("terminal length 0\n"))
	input.Write([]byte(cmd.Text))

	bytes := make([]byte, 0, 1024)
	for {
		select {
		case l := <-newLine:
			bytes = append(bytes, l...)
		case <-time.After(time.Millisecond * 100):
			session.Close()
			return bytes, err
		}
	}
}
