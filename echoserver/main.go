package main

import (
	"net"
	"fmt"
	"os"
	"strings"
)

func main() {
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Client error %s", err.Error())
			continue
		}
		go handleClient(conn)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Client error %s", err.Error())
			continue
		}
	}
}

func handleClient(conn net.Conn) error {
	var buf [512]byte
	for {
		n, err := conn.Read(buf[:])
		if err != nil {
			return err
		}
		echo := buf[:n]
		secho := strings.TrimSpace(string(echo))
		if secho == "quit" {
			conn.Close()
			return nil
		}
		fmt.Fprintln(os.Stderr, secho)
		_, err = conn.Write(echo)
		if err != nil {
			return err
		}
	}
	panic("unreachable")
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
