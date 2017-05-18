package Wizard_Text

import (
	"fmt"
	"net"
	"os"
	"strings"
	"bufio"
	"log"
	"github.com/bshafer93/WizardText"
)

const (
	CONN_HOST = "107.170.196.189"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)


func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	nullCount := 0
	connActive := true





	for connActive == true {

		msg, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		stringedMsg := string(msg)
		headerMsg := strings.HasPrefix(stringedMsg, "MIncoming")

		if headerMsg == true {
			if nullCount == 0 {
				nullCount = 0
			}
			nullCount--
			//fmt.Print("Header Received")
		}

		if nullCount <= 5 && headerMsg == false {



			if len(stringedMsg) != 0 {
				r := strings.NewReplacer("<", "&lt",
					">", "&gt",
					"&","&amp")
				sanitized := r.Replace(stringedMsg)
				fmt.Print("Message Receivede:", string(msg))
				conn.Write([]byte(sanitized + "\n"))
			}
			nullCount++

		}


	if nullCount >= 5 {
		connActive = false
	}
}

	conn.Close()
}



