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
	name string
	power int
	cost int
}

type ChatMSG struct {
	IncomingMSG
	msg string

}

type MSG interface {
	deduceCommand() string
	deduceContents() string
	sanitizeMessage() string
}

func (I *IncomingMSG) deduceCommand() string{


	stringedMsg := I.content

	switch {
	case strings.HasPrefix(stringedMsg, "heartbeat"):
		I.whatType = "heartbeat"
		return I.whatType
	case strings.HasPrefix(stringedMsg, "/"):
		I.whatType = "Command"
		return I.whatType
	case strings.HasPrefix(stringedMsg, "@"):
		I.whatType = "Invite"
		return I.whatType
	default:
		I.whatType = "Simple_Message"
		return I.whatType
	}

}

func (I *IncomingMSG) deduceContent() string {
	msg, err := bufio.NewReader(I.conn).ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	I.content = string(msg)
	return I.content

}

func (I *IncomingMSG) sanitizeMessage() string {
	msg, err := bufio.NewReader(I.conn).ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	I.content = string(msg)
	return I.content

}






