package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

var src string

func read(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	return string(b)
}

func uploadFile(connection *ftp.ServerConn) {
	fmt.Scan(&src)
	data := read(src)
	err := connection.Stor(src, bytes.NewBufferString(data))
	fmt.Print("file is save")
	if err != nil {
		log.Fatal(err)
	}
}

func downloadFile(connection *ftp.ServerConn) {
	fmt.Scan(&src)
	res, err := connection.Retr(src)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()

	outFile, err := os.Create(src)
	if err != nil {
		log.Fatal(err)
	}

	defer outFile.Close()

	_, err = io.Copy(outFile, res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("file is download")
}

func createDir(connection *ftp.ServerConn) {
	fmt.Scan(&src)
	err := connection.MakeDir(src)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("create")
}

func deleteFile(connection *ftp.ServerConn) {
	fmt.Scan(&src)
	err := connection.Delete(src)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("delete")
}

func removeDir(connection *ftp.ServerConn) {
	fmt.Scan(&src)
	err := connection.RemoveDir(src)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Dir is deleted")
}

func listDir(connection *ftp.ServerConn) {
	src, err := connection.CurrentDir()
	if err != nil {
		log.Fatal(err)
	}
	entries, err := connection.NameList(src)
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range entries {
		fmt.Println(s)
	}
}

func main() {
	connection, err := ftp.Dial("students.yss.su", ftp.DialWithTimeout(5*time.Second))
	// connection, err := ftp.Dial("185.20.227.83:13370", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Quit()

	err = connection.Login("ftpiu8", "3Ru7yOTA")
	// err = connection.Login("amaguma", "123qwe")
	if err != nil {
		log.Fatal(err)
	}

	// err = connection.ChangeDir("Umar")
	if err != nil {
		log.Fatal(err)
	}

	var comand string
	for {
		fmt.Scan(&comand)
		switch comand {
		case "upload":
			uploadFile(connection)
		case "download":
			downloadFile(connection)
		case "mkdir":
			createDir(connection)
		case "del":
			deleteFile(connection)
		case "ls":
			listDir(connection)
		case "removeDir":
			removeDir(connection)
		default:
			fmt.Print("The command is not recognized")
		}
	}
}
