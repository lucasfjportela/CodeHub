package main

import (
	"codehub-sd/messageFormat"
	"encoding/gob"
	"net"
	"os"
	"os/exec"
	"fmt"

	"github.com/jlaffaye/ftp"
)

/*
type Req struct {
	Login      string
	Password   string
	Authorized bool
}

*/

//type ListUsers [] messageFormat.MessageFormat

func handleClientDNSConnection(conn *net.TCPConn, msg messageFormat.MessageFormat) string {

	dnsResponse := &messageFormat.MessageFormat{}

	encoder := gob.NewEncoder(conn)

	encoder.Encode(msg)

	//readed := make([]byte, 1024)
	//datasize, _ := conn.Read(readed)
	//data := readed[:datasize]

	decoder := gob.NewDecoder(conn)

	decoder.Decode(dnsResponse)
	
	return dnsResponse.Payload.(string)
}

func handleClientAuthConnection(conn *net.TCPConn, msgUser messageFormat.MessageFormat) bool {

	encoderServer := gob.NewEncoder(conn)
	encoderServer.Encode(msgUser)

	var authResponse bool

	decoder := gob.NewDecoder(conn)

	decoder.Decode(&authResponse)

	//readed := make([]byte, 1024)
	//datasize, _ := conn.Read(readed)
	//data := readed[:datasize]
	return authResponse
}

/*
func handleClientServerConnection(conn *net.TCPConn, msgUser messageFormat.MessageFormat) {
	defer conn.Close()

	msg := msgUser

	encoder := gob.NewEncoder(conn)
	encoder.Encode(msg)

}
*/

func handleServerConnection(reqType string, serverAddr string) {

	data := make([]byte, 1024)
	conn, err := ftp.Dial(serverAddr)
	conn.Login("admin", "admin")
	loginErr := conn.Login("admin", "admin")

	if loginErr != err {
		panic(loginErr)
	}
	if err != nil {
		panic(err)
	}

	if reqType == "str" {
		file, err := os.Open("Wedson.txt")
		if err != nil {
			panic(err)
		}
		conn.Stor("/User/Theu.txt", file)
	}

	if reqType == "get" {
		test, _ := conn.Retr("MATHEUS.txt")

		file, _ := os.Create("Wedson.txt")
		test.Read(data)
		file.Write(data)
	}

}
func main() {
	
	// Console clear
	cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
	cmd.Run()
	
	fmt.Println("------------------ Codehub ------------------\n\n")
	fmt.Println("")

	var userAction, login, password, filename string

	fmt.Scanln(&userAction, &login, &password)
	//fmt.Println(userAction, login, password)
	//fmt.Printf("%s", action)

	if(userAction != "auth"){
		panic("\nAuthentication is needed!\nTry:\n auth <login> <password>\n\n")
	}

	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:2223")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	msg := messageFormat.MessageFormat{
		Origin:  "Client",
		ReqType: "Auth",
	}
	authAddr := handleClientDNSConnection(conn, msg)

	msg = messageFormat.MessageFormat{
		Payload: []string{login, password},
	}

	tcpAuth, _ := net.ResolveTCPAddr("tcp", authAddr)
	connAuth, _ := net.DialTCP("tcp", nil, tcpAuth)

	authResponse := handleClientAuthConnection(connAuth, msg)

	if !authResponse {
		panic("Unauthorized")
	}

	msg = messageFormat.MessageFormat{
		Origin:  "Client",
		ReqType: "Server",
	}
	fmt.Println("oiau")
	serverAddr := handleClientDNSConnection(connAuth, msg)
	fmt.Println("oi")

	// Client get/str files
	fmt.Println("\nuse get name-of-file.txt to get a code from server\n")
	fmt.Println("\nuse str name-of-file.txt to upload a code to server\n")
	fmt.Scanln(&userAction, &filename)

	if(userAction != "get" || userAction != "str"){
		panic("\nWrong command, try again.\n\n")
	}

	if(userAction == "get"){
		handleServerConnection("get", serverAddr)
	} else {
		handleServerConnection("str", serverAddr)
	}
}
