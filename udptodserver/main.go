package main

import (
	"net"
	"fmt"
	"os"
	"time"
)

func main() {
	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("up4", service)
	checkError(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	for {
		var buf [512]byte
		_, addr, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Client connection error %s", err.Error())
			continue
		}
		daytime := time.Now().String()
		_, err = conn.WriteToUDP([]byte(daytime+"\n"), addr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Client write error %s", err.Error())
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
