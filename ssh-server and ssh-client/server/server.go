package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gliderlabs/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func sessionHandler(s ssh.Session) {
	args := s.Command()
	if len(args) > 0 {
		if len(args) < 2 {
			io.WriteString(s, "path required\n")
			return
		}
		switch args[0] {
		case "ls":
			f, err := ioutil.ReadDir(args[1])
			if err != nil {
				io.WriteString(s, err.Error()+"\n")
				break
			}
			for _, v := range f {
				io.WriteString(s, v.Name()+"\n")
			}
			break
		case "mkdir":
			err := os.Mkdir(args[1], os.ModePerm)
			if err != nil {
				io.WriteString(s, err.Error()+"\n")
			}
			break
		case "rmdir":
			err := os.Remove(args[1])
			if err != nil {
				io.WriteString(s, err.Error()+"\n")
			}
			break
		default:
			io.WriteString(s, "command not found\n")
		}

		return
	}

	io.WriteString(s, "access success\n")
	term := terminal.NewTerminal(s, "> ")
	for {
		d, err := term.ReadLine()
		if err != nil {
			if err != io.EOF {
				fmt.Println(err.Error())
			}
			break
		}

		args := strings.Fields(string(d))
		if len(args) < 1 {
			io.WriteString(s, "command required\n")
			continue
		}

		if args[0] == "exit" {
			break
		}

		if len(args) < 2 {
			io.WriteString(s, "path required\n")
			continue
		}

		switch args[0] {
		case "ls":
			f, err := ioutil.ReadDir(args[1])
			if err != nil {
				io.WriteString(s, err.Error()+"\n")
				break
			}
			for _, v := range f {
				io.WriteString(s, v.Name()+"\n")
			}
			break
		case "mkdir":
			err := os.Mkdir(args[1], os.ModePerm)
			if err != nil {
				io.WriteString(s, err.Error()+"\n")
			}
			break
		case "rmdir":
			err := os.Remove(args[1])
			if err != nil {
				io.WriteString(s, err.Error()+"\n")
			}
			break
		default:
			io.WriteString(s, "command not found\n")
		}
	}
}

func passwordHandler(ctx ssh.Context, password string) bool {
	return ctx.User() == "qwerty" && password == "bydfkbl5151"
}

func main() {
	s := &ssh.Server{
		Addr:            ":2209",
		Handler:         sessionHandler,
		PasswordHandler: passwordHandler,
	}

	err := s.SetOption(ssh.HostKeyFile("id_rsa"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print("kek")
	s.ListenAndServe()
}
