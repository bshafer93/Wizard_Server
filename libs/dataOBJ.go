package libs

import (
	"net"
	"bufio"
	"strings"
	"log"
)

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
		return I.WhatType
	}

}

func (I *IncomingMSG) DeduceContent() string {
	msg, err := bufio.NewReader(I.Conn).ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	I.Content = string(msg)
	return I.Content

}

func (I *IncomingMSG) SanitizeMessage() string {
	msg, err := bufio.NewReader(I.Conn).ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	I.Content = string(msg)
	return I.Content

}






