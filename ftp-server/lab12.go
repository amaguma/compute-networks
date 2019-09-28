package main

import (
	"flag"
	"log"

	filedriver "github.com/goftp/file-driver"
	"github.com/goftp/server"
)

func main() {
	var (
		root = flag.String("root", "/home/iu9_32_09/ftproot", "Root directory to serve")
		user = flag.String("user", "amaguma", "Username for login")
		pass = flag.String("pass", "123qwe", "Password for login")
		port = flag.Int("port", 13370, "Port")
		host = flag.String("host", "185.20.227.83", "Host")
	)
	flag.Parse()

	factory := &filedriver.FileDriverFactory{
		RootPath: *root,
		Perm:     server.NewSimplePerm("iu9_32_09", "iu9_32_09"),
	}

	opts := &server.ServerOpts{
		Factory:  factory,
		Port:     *port,
		Hostname: *host,
		Auth:     &server.SimpleAuth{Name: *user, Password: *pass},
	}

	log.Printf("Starting ftp server on %v:%v", opts.Hostname, opts.Port)
	log.Printf("Username %v, Password %v", *user, *pass)
	server := server.NewServer(opts)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
