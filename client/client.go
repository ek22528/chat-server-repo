package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func read(conn net.Conn) {
	//TODO In a continuous loop, read a message from the server and display it.
  for{
    buf := make([]byte, 1024)
    length, err := conn.Read(buf)
    if err != nil {
      fmt.Println("client error:  ", err.Error())
      break
    }
    fmt.Println(string(buf[:length]))
  }
}

func write(conn net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
  for {
    var chat string
    fmt.Scanln(&chat)
    _, err := conn.Write([]byte(chat))
    if err != nil {
      fmt.Println("client error:  ", err.Error())
      break 
    }
  }
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
  fmt.Println(*addrPtr)
	//TODO Try to connect to the server
  conn, err := net.Dial("tcp", *addrPtr)
  if err != nil {
    fmt.Println("client error:  ", err.Error())
    os.Exit(1)
    return
  }
	//TODO Start asynchronously reading and displaying messages
  go read(conn)
  go write(conn)
	//TODO Start getting and sending user messages.
  time.Sleep(time.Hour)
}
