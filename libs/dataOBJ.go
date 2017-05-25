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
	_ "github.com/go-sql-driver/mysql"
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

type UserCheck struct {
	Auth
	UserReg
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
	case strings.HasPrefix(stringedMsg, "/Register"):
		I.WhatType = "UserReg"
		return I.WhatType
	case strings.HasPrefix(stringedMsg, "/Login"):
		I.WhatType = "Login"
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
		 log.Print("Uh-Oh!: ",err)
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

func NewUserReg(Username string, Password string,UserEmail string ) *UserReg {
	sr := UserReg{
		Username: Username,
		Password: Hashpass(Password)  ,
		Email: UserEmail,
	}

	sr.Register()
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

	db := OpenDB()
	stmt, err := db.Prepare("INSERT INTO login VALUES (?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(I.Username,I.Password,I.Email)
	fmt.Println("New user Registered!")
	PrintLoginPeeps()
	db.Close()


}
func  Login(){

	db := OpenDB()

	var user UserCheck
	row, err := db.Query("SELECT username,email FROM login WHERE username=?", "Brent")
	fmt.Println(row,"This is a row")
	if err != nil {
		log.Fatal(err)
	}
	for row.Next() {
		errr := row.Scan(&user.Username, &user.Email)
		if errr != nil {
			log.Fatal(errr)
		}
		log.Println(user.Username, user.Email)
	}

	fmt.Println(row)

	db.Close()


}




func OpenDB() *sql.DB {
	db, err := sql.Open("mysql", "root:longleaf1@tcp(107.170.196.189:3306)/users")
	if err != nil {
		db.Close()
		panic(err)
		log.Fatal(err)
	}
	fmt.Println("Connected!")
	return db
}

func  PrintLoginPeeps(){
	db := OpenDB()
	// Execute the query
	rows, err := db.Query("SELECT * FROM login")
	if err != nil {
		panic(err.Error())
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	// Make a slice for the values
	values := make([]interface{}, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		// Print data
		for i, value := range values {
			switch value.(type) {
			case nil:
				fmt.Println(columns[i], ": NULL")

			case []byte:
				fmt.Println(columns[i], ": ", string(value.([]byte)))

			default:
				fmt.Println(columns[i], ": ", value)
			}
		}
		fmt.Println("-----------------------------------")
	}

}
