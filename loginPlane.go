package main

import (
	"fmt"
	"net"
	"os"
	"github.com/bshafer93/Wizard_Server/libs"
	"runtime"
	"time"
	"io"
)

const (
	//CONN_HOST = "107.170.196.189"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)


func main() {
	CONN_HOST := "107.170.196.189"
	if runtime.GOOS == "windows" {
		CONN_HOST = "192.168.0.25"
	}
	Lobby := libs.NewServerRoom()
	//Lobby.Broadcast := make(chan string)
	//Lobby.Receive := make(chan string)
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		 fmt.Println("Error listening:", err.Error())
		 os.Exit(1)
	}
	// Close the listener when the application closes.
	//defer l.Close()

	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		Lobby.UserList = make(map[string]net.Conn)
		Lobby.UserList["Player1"] = conn
		go handleRequest(Lobby.UserList["Player1"])


	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {

	nullCount := 0
	connActive := true
	for connActive == true {

		one := []byte{}
		conn.SetReadDeadline(time.Now())
		if _, err := conn.Read(one); err == io.EOF {
			conn.Close()
			conn = nil
		}


		var content libs.IncomingMSG
		content.Conn = conn
		content.Content = content.DeduceContent()
		content.WhatType = content.DeduceCommand()


		if content.WhatType == "heartbeat" {
			if nullCount == 0 {
				nullCount = 0
			}
			nullCount--
		}

		if content.Content != "Sent_Nothing" && content.WhatType == "Simple_Message" {
			content.SendToAll()
			nullCount++
		}


	if nullCount >= 5 {
		connActive = false
	}

}

	conn.Close()
}



