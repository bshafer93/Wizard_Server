package libs

import (
	"net"
	"bufio"
	"strings"
	"log"
	"fmt"
	"math/rand"
)

type UserConn struct {
	Conn net.Conn
	Username string
	Authorized bool

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
	default:
		I.WhatType = "Simple_Message"
		I.Content = SanitizeMessage(I.Content)
		return I.WhatType
	}

}

func (I *IncomingMSG) DeduceContent() string {
	msg, err := bufio.NewReader(I.Conn).ReadString('\n')

	if err != nil {
		log.Print(err)

	}

	I.Content = string(msg)
	return I.Content

}

func SanitizeMessage(s string) string {

	if len(s) != 0 {
		r  := strings.NewReplacer("<", "&lt",
			">", "&gt",
			"&","&amp")

		sanitized := r.Replace(s)
		return sanitized
	} else {return "Sent_Nothing" }

}

func (I *IncomingMSG) SendToAll() {

	San := SanitizeMessage(I.Content)
	fmt.Printf(San)
	I.Conn.Write([]byte(San + "\n"))
	//if err != nil{
	//	fmt.Println("Error Sending Message:", err.Error())
		//I.Conn.Close() // Closes Connection

	//}


}

func NewServerRoom() *ServerRoom{
	randNum := rand.Int()


	sr := ServerRoom{
		Name: string(randNum),
		ID: randNum,

	}
	return &sr
}






