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
	"strconv"
	"strings"
)

const (
	//CONN_HOST = "107.170.245.15"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"

)

var Lobby = libs.NewServerRoom()

func main() {

	Lobby.UserList = make(map[libs.User.username]libs.User.Conn)
	userInt := 0
	cert, err := tls.LoadX509KeyPair("certs/server.pem", "certs/server.key")

	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}


	config := tls.Config{Certificates: []tls.Certificate{cert}}
	config.Rand = rand.Reader


	CONN_HOST := "127.0.0.1"

	if runtime.GOOS == "windows" {
		CONN_HOST = "192.168.0.25"
	}

	//Lobby.Receive := make(chan string)
	// Listen for incoming connections.
	l, err := tls.Listen(CONN_TYPE, CONN_HOST+ ":" + CONN_PORT,&config)
	if err != nil {
		 fmt.Println("Error listening:", err.Error())
		 os.Exit(1)
	}
	// Close the listener when the application closes.
	//defer l.Close()

	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		userInt++
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



		go handleRequest(conn,Lobby)
	}
}




// Handles incoming requests.
func handleRequest(conn net.Conn) {
	var connUser libs.User
	connUser.Conn = conn
		// Global var for conection status
	connActive := true

	// Essentially the update loop for the connection
	for connActive == true {

		IncomingMessage := libs.NewIncomingMSG(conn)

		/* ######################### HANDLE INCOMING CONTENT ######################### */
			libs.HandleIncomingMessage(IncomingMessage,&Lobby)	
		/* ######################### END HANDLE INCOMING CONTENT ######################### */


	fmt.Printf("I Disconnected")
	delete(Lobby.UserList,connUser.Username)
	conn.Close()
}



