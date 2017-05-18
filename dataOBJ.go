package Wizard_Text

import (
	"net"
	"bufio"
	"strings"
	"log"
)

type IncomingMSG struct{
	net.Conn
	whatType string
	content string
}
type Spell struct {
	IncomingMSG
	name string
	power int
	cost int
}

type ChatMSG struct {
	IncomingMSG

}

type MSG interface {
	deduceCommand() string
	deduceContents() string
}

func (I *IncomingMSG) deduceCommand() string{


	stringedMsg := I.content

	switch {
	case strings.HasPrefix(stringedMsg, "heartbeat"):
		I.whatType = "Heartbeat"
		return I.whatType
	case strings.HasPrefix(stringedMsg, "/"):
		I.whatType = "Command"
		return I.whatType
	case strings.HasPrefix(stringedMsg, "@"):
		I.whatType = "Invite"
		return I.whatType
	default:
		I.whatType = "Simple Message"
		return I.whatType
	}

}

func (I *IncomingMSG) deduceContent() string {
	msg, err := bufio.NewReader(I).ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	I.content = string(msg)
	return I.content

}





