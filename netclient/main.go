package main

import (
	"flag"
	"net"
	"os"
)

var tcp = flag.Bool("tcp", false, "Use TCP to connect")
var udp = flag.Bool("udp", false, "Use UDP to connect")
var addr = flag.String("a", "", "Sets the remote address ip:port to send messages to")
var msg = flag.String("m", "", "The message to send")

func main() {
	flag.Parse()
	switch {
	case *tcp && *udp:
		println("You can't use the -tcp and -udp flags at the same time.")
		os.Exit(1)
	case !*tcp && !*udp:
		doTcp()
	case *tcp:
		doTcp()
	case *udp:
		doUdp()
	}
}

func doUdp() {
	udpAddr, err := net.ResolveUDPAddr("udp", *addr)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	_, err = conn.Write([]byte(*msg))
	checkError(err)
	var buf [512]byte
	n, err := conn.Read(buf[:])
	checkError(err)
	println(string(buf[:n]))
}

func doTcp() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", *addr)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	checkError(err)
	_, err = conn.Write([]byte(*msg))
	checkError(err)
	var buf [512]byte
	n, err := conn.Read(buf[:])
	checkError(err)
	println(string(buf[:n]))
}

func checkError(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
