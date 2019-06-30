package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const addr = "0.0.0.0:8888"

func main() {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		fmt.Println("ResolveTCPAddr:", err.Error())
		return
	}

	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		fmt.Println("DialTCP:", err.Error())
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprint(conn, text+"\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("->: " + message)

		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			conn.Close()
			return
		}
	}
}
