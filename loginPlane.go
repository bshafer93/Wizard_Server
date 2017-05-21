package main

import (
	"fmt"
	"os"
	"github.com/bshafer93/Wizard_Server/libs"
	"runtime"
	"crypto/tls"
	"log"
	"crypto/rand"
	"crypto/x509"
	"net"
)

const (
	//CONN_HOST = "107.170.196.189"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)


func main() {

	cert, err := tls.LoadX509KeyPair("certs/server.pem", "certs/server.key")
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	config.Rand = rand.Reader


	CONN_HOST := "107.170.196.189"
	if runtime.GOOS == "windows" {
		CONN_HOST = "192.168.0.25"
	}
	Lobby := libs.NewServerRoom()
	//Lobby.Broadcast := make(chan string)
	//Lobby.Receive := make(chan string)
	// Listen for incoming connections.
	l, err := tls.Listen(CONN_TYPE, CONN_HOST + ":" + CONN_PORT,&config)
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


		tlscon, ok := conn.(*tls.Conn)
		if ok {
			log.Print("ok=true")
			state := tlscon.ConnectionState()
			for _, v := range state.PeerCertificates {
				log.Print(x509.MarshalPKIXPublicKey(v.PublicKey))
			}
			// Handle connections in a new goroutine.

		}

		Lobby.UserList = make(map[string]net.Conn)
		Lobby.UserList["Player1"] = conn
		go handleRequest(Lobby.UserList["Player1"])
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {

	connActive := true
	for connActive == true {

		content := libs.NewIncomingMSG(conn)

		if content.Content == "Client Disconnected"{
			//If client is gone, disconnect and end loop
			connActive = false
			return

		}

		if content.Content != "Sent_Nothing" && content.WhatType == "Simple_Message" {
			content.SendToAll()
		}




}
	fmt.Printf("I Disconnected")
	conn.Close()
}



