//client.go:
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Подключаемся к сокету
	conn, _ := net.Dial("tcp", "185.20.227.83:8010")
	for {
		// Чтение входных данных от stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Vectors values (delimited by space): ")
		text, _ := reader.ReadString('\n')
		// Отправляем в socket
		fmt.Fprintf(conn, text+"\n")
		// Прослушиваем ответ
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print(message)
	}
}
