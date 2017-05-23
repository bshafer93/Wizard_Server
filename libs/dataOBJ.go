package libs

import (
	"bufio"
	"strings"
	"log"
	"fmt"
	"math/rand"
	"net"
	"html"
	"golang.org/x/crypto/bcrypt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

const (
	DB_HOST = "tcp(127.0.0.1:3306)"
	DB_NAME = "users"
	DB_USER = "root"
	DB_PASS = "longleaf1"

)

type UserConn struct {
	Conn net.Conn
	Username string
	Authorized bool

}

type UserReg struct{
	Username string
	Password string
	Email string
	Auth
}

type SpellBook struct {
	Spellbook []Spell
}

type ServerRoom struct {
	Name string
	ID int
	UserList map[string]net.Conn
	Broadcast chan int
	Receive chan int
}

type IncomingMSG struct{
	Conn net.Conn
	WhatType string
	Content string
}
type Spell struct {
	Name string
	Power int
	Cost int
	Description string
}

type ChatMSG struct {
	IncomingMSG
	Msg string

}


type Auth interface {
	Register()
	HashPass()
	StorePass()
	Login()
}


type MSG interface {
	DeduceCommand() string
	DeduceContents() string
	SanitizeMessage() string
	SendToAll()
}

type Book interface {
	alphabatize()
}

func (I *IncomingMSG) DeduceCommand() string{


	stringedMsg := I.Content

	switch {
	case strings.HasPrefix(stringedMsg, "heartbeat"):
		I.WhatType = "heartbeat"
		return I.WhatType
	case strings.HasPrefix(stringedMsg, "/"):
		I.WhatType = "Command"
		return I.WhatType
	case strings.HasPrefix(stringedMsg, "@"):
		I.WhatType = "Invite"
		return I.WhatType
	case strings.HasPrefix(stringedMsg, "/Register"):
		// Run Register Fuctions here
		// What would you like your user name to be?
		// Password?
		//Email
		//Ping back User and Email Check if yes or no
		I.WhatType = "UserReg"
		return I.WhatType
	default:
		I.WhatType = "Simple_Message"
		I.Content = SanitizeMessage(I.Content)
		return I.WhatType
	}

}

func (I *IncomingMSG) DeduceContent() string {
	msg, err := bufio.NewReader(I.Conn).ReadString('\n')

	if err != nil {
		 I.Conn.Close()
		 log.Print("Fuck ",err)
		// If client disconnects tell server
		return "Client Disconnected"


	}


	I.Content = string(msg)
	return I.Content

}

func SanitizeMessage(s string) string {

	if len(s) != 0 {

		r := html.EscapeString(s)



		return r
	} else {return "Sent_Nothing" }

}

func (I *IncomingMSG) SendToAll() {

	San := SanitizeMessage(I.Content)

	_,errr := fmt.Printf(San)
	if errr != nil{
		fmt.Println("Error Sending Message:", errr.Error())
		I.Conn.Close() // Closes Connection

	}

	_, err := I.Conn.Write([]byte(San + "\n"))

	if err != nil{
		fmt.Println("Error Sending Message:", err.Error())
		I.Conn.Close() // Closes Connection

	}


}

func NewServerRoom() *ServerRoom{
	randNum := rand.Int()


	sr := ServerRoom{
		Name: string(randNum),
		ID: randNum,

	}
	return &sr
}

func NewIncomingMSG(conn net.Conn) *IncomingMSG {
	IC := new(IncomingMSG)
	IC.Conn = conn
	IC.Content = IC.DeduceContent()
	IC.WhatType = IC.DeduceCommand()
	return IC

}

func ServerPrivateMessage(c net.Conn,s string){


	_, err := c.Write([]byte(s))

	if err != nil{
		fmt.Println("Error Sending Message:", err.Error())
		c.Close() // Closes Connection

	}

}

func NewUserReg(Username string, Password string,Email string ) *UserReg {
	sr := UserReg{
		Username: Username,
		Password: Hashpass(Password)  ,
		Email: "Mynameisnicky@gmail.com",

	}
	return &sr

}

func Hashpass(pass string) string {




		hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}

	fmt.Println("Hash to store:", string(hash))
	return string(hash)
}





func (I *UserReg) Register(){

}




func OpenDB() *sql.DB {
	db, err := sql.Open("mymysql", fmt.Sprintf("%s/%s/%s", DB_NAME, DB_USER, DB_PASS))
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	fmt.Println("Connected!")
	return db
}

