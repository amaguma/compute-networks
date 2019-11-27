//server.go:
package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Launching server...")
	// Устанавливаем прослушивание порта
	ln, _ := net.Listen("tcp", ":8010")
	// Открываем порт
	conn, _ := ln.Accept()
	for {
		// Будем прослушивать все сообщения разделенные \n
		message, _ := bufio.NewReader(conn).ReadString('\n')
		values := strings.Fields(message)
		fmt.Println("Parsing vectors (" + values[0] + "; " + values[1] + "; " + values[2] + "), (" + values[3] + "; " + values[4] + "; " + values[5] + ")")
		// переводим координаты в тип float
		x1, _ := strconv.ParseFloat(values[0], 10)
		y1, _ := strconv.ParseFloat(values[1], 10)
		z1, _ := strconv.ParseFloat(values[2], 10)
		x2, _ := strconv.ParseFloat(values[3], 10)
		y2, _ := strconv.ParseFloat(values[4], 10)
		z2, _ := strconv.ParseFloat(values[5], 10)
		conn.Write([]byte(fmt.Sprintf("%f\n", x1*x2+y1*y2+z1*z2)))
	}
}
