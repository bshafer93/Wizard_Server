package libs

import (
	"net"
	"bufio"
	"strings"
	"log"
)

type UserConn struct {
	Conn net.Conn
}

type IncomingMSG struct{
	Conn net.Conn
	WhatType string
	Content string
}
type Spell struct {
	IncomingMSG
	Name string
	Power int
	Cost int
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
	I.Conn.Write([]byte(San + "\n"))


}








