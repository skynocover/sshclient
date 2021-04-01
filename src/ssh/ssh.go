package ssh

import (
	"log"

	"github.com/helloyi/go-sshclient"
	"golang.org/x/crypto/ssh"
)

func Dial(host, user, password string) {
	client, err := sshclient.DialWithPasswd(host, user, password)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	config := &sshclient.TerminalConfig{
		Term:   "xterm",
		Height: 40,
		Weight: 80,
		Modes: ssh.TerminalModes{
			ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
			ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
		},
	}

	if err := client.Terminal(config).Start(); err != nil {
		log.Fatalln(err)
	}
}
