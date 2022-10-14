package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func read(conn net.Conn) {
	//TODO In a continuous loop, read a message from the server and display it.
	reader := bufio.NewReader(conn)
	for {
		newMsg, _ := reader.ReadString('\n')
		fmt.Print(newMsg)
	}
}

func write(conn net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
	reader := bufio.NewReader(os.Stdin)
	newMsg, _ := reader.ReadString('\n')
	if newMsg != "\n" {
		fmt.Fprintln(conn, newMsg)
	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	fmt.Println("recieved flag: ", *addrPtr)
	//TODO Try to connect to the server

	conn, _ := net.Dial("tcp", *addrPtr)
	fmt.Println("Success! - connected to ", *addrPtr)

	//TODO Start asynchronously reading and displaying messages
	go read(conn)
	//TODO Start getting and sending user messages.
	for {
		write(conn)
	}

}
