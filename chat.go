package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type Client chan<- string

var (
	incommingClients = make(chan Client)
	leavingClients   = make(chan Client)
	chatMessages     = make(chan string)
	host             = flag.String("host", "localhost", "host")
	port             = flag.Int("port", 3090, "port")
)

func HandleConn(conn net.Conn) {
	defer conn.Close()
	clientMessages := make(chan string)
	go MessageWriter(conn, clientMessages)
	clientName := conn.RemoteAddr().String()
	clientMessages <- fmt.Sprintf("Welcome to the server, your name is: %s\n", clientName)
	incommingClients <- clientMessages
	chatMessages <- fmt.Sprintf("%s has joined the server\n", clientName)

	input := bufio.NewScanner(conn)
	for input.Scan() {
		chatMessages <- fmt.Sprintf("%s: %s\n", clientName, input.Text())
	}

	leavingClients <- clientMessages
	chatMessages <- fmt.Sprintf("%s has left the server\n", clientName)
}

func MessageWriter(conn net.Conn, messages <-chan string) {
	for msg := range messages {
		fmt.Fprintln(conn, msg)
	}
}

func Broadcast() {
	clients := make(map[Client]bool)
	for {
		select {
		case msg := <-chatMessages:
			for client := range clients {
				client <- msg
			}

		case newClient := <-incommingClients:
			clients[newClient] = true

		case leavingClient := <-leavingClients:
			delete(clients, leavingClient)
			close(leavingClient)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}
	go Broadcast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go HandleConn(conn)
	}
}
