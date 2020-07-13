package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("command required")
		return
	}

	config := &ssh.ClientConfig{
		User: "qwerty",
		Auth: []ssh.AuthMethod{
			ssh.Password("bydfkbl5151"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", "185.20.227.83:2209", config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	session, err := client.NewSession()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	b, err := session.CombinedOutput(strings.Join(os.Args[1:], " "))
	fmt.Print(string(b))
	session.Close()
}
