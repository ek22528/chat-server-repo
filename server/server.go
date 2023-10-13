package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	// TODO: all
	// Deal with an error event.
  if err != nil {
    fmt.Println("Server Error:   ", err.Error())
    os.Exit(1)
  }
  
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	// TODO: all
	// Continuously accept a network connection from the Listener
	// and add it to the channel for handling connections.
  for {
    conn, err := ln.Accept() 
    handleError(err)
    conns <- conn
  }
}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	// TODO: all
	// So long as this connection is alive:
	// Read in new messages as delimited by '\n's
	// Tidy up each message and add it to the messages channel,
	// recording which client it came from.
  for {
    buf := make([]byte, 1024)
    length, err := client.Read(buf)
    if err != nil {
      return
    }
    handleError(err)
    msg := Message{
      sender: clientid,
      message: string(buf[:length]),
    }
    msgs <- msg
  }
}

func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()

  listenAddress := "127.0.0.1" + string(*portPtr)

	//TODO Create a Listener for TCP connections on the port given above.
  ln, err := net.Listen("tcp", listenAddress)
  handleError(err)

	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)

	//Start accepting connections
	go acceptConns(ln, conns)
  lastid := 0
	for {
		select {
		case conn := <-conns:
			//TODO Deal with a new connection
			// - assign a client ID
			// - add the client to the clients channel
			// - start to asynchronously handle messages from this client
      clients[lastid] = conn
      go handleClient(conn, lastid, msgs)
      lastid++
      fmt.Println("new client ", lastid)
		case msg := <-msgs:
			//TODO Deal with a new message
			// Send the message to all clients that aren't the sender
      for i, client := range clients {
        if i == msg.sender {
          continue
        }
        _, err := client.Write([]byte(msg.message))
        handleError(err)
      }
		}
	}
}
