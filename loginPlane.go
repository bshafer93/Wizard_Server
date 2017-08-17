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
)

const (
	//CONN_HOST = "107.170.196.189"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"

)

var Lobby = libs.NewServerRoom()

func main() {
	Lobby.UserList = make(map[string]net.Conn)

	userInt := 0
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
func handleRequest(conn net.Conn, Lobby *libs.ServerRoom) {
	var connUser libs.User

	connUser.Conn = conn
	connActive := true
	for connActive == true {

		content := libs.NewIncomingMSG(conn)

		if content.Content == "Client Disconnected"{
			//If client is gone, disconnect and end loop
			delete(Lobby.UserList,connUser.Username)
			connActive = false
			return

		}


		switch content.WhatType {

		case "UserReg":
			fmt.Println("Got here!")
			libs.ServerPrivateMessage(content.Conn,"What would you like your user name to be?")
			UsernameConn := libs.NewIncomingMSG(conn)
			if libs.CheckUsername(UsernameConn.Content) == false {
				libs.ServerPrivateMessage(content.Conn, "What would you like your password to be?")
				Pwd := libs.NewIncomingMSG(conn)
				libs.ServerPrivateMessage(content.Conn, "What would you like your email to be?")
				email := libs.NewIncomingMSG(conn)
				libs.NewUserReg(UsernameConn.Content, Pwd.Content, email.Content)
				libs.ServerPrivateMessage(content.Conn, "Now registered!")
			} else {
				libs.ServerPrivateMessage(content.Conn,"Username Taken! Try another.")
			}

		case "Login":
			libs.ServerPrivateMessage(content.Conn,"What is your username?")
			Username := libs.NewIncomingMSG(conn)
			libs.ServerPrivateMessage(content.Conn,"What is your password?")
			Pwd := libs.NewIncomingMSG(conn)
			connUser.Username = content.Login(Username.Content,Pwd.Content)
			if connUser.Username == Username.Content{
				Lobby.UserList[connUser.Username] = conn
				PH := libs.RetrieveHealth(connUser.Username)
				PM := libs.RetrieveMana(connUser.Username)
				PL := libs.RetrieveMana(connUser.Username)
				fmt.Println((strconv.Itoa(PH) +strconv.Itoa(PM) ))
				libs.ServerPrivateMessage(content.Conn,"PH"+strconv.Itoa(PH))
				libs.ServerPrivateMessage(content.Conn,"PM"+strconv.Itoa(PM))
				libs.ServerPrivateMessage(content.Conn,"PL"+strconv.Itoa(PL))
				fmt.Println(connUser.Username+">Has Connected!")

			}else{


			}


		case "adminCommand":
			libs.ServerPrivateMessage(content.Conn,connUser.Username+">The fuck you want?")

		case "Spell":
				RemoveHash := content.Content[0:len(content.Content)]
			switch RemoveHash {

			case "Fireball":
				libs.ServerPrivateMessage(content.Conn,"Recipient?")
				R := libs.NewIncomingMSG(conn)
				libs.Fireball(connUser,R.Content,Lobby.UserList)

			}


		case "Simple_Message":
			if len(connUser.Username) == 0{
				libs.ServerPrivateMessage(content.Conn,"Server> Go logon!")

			} else {
				content.SendToAll(connUser.Username,Lobby.UserList)
			}
		}







}
	fmt.Printf("I Disconnected")
	delete(Lobby.UserList,connUser.Username)
	conn.Close()
}



