package main

import (
	"codehub-sd/messageFormat"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"os/exec"
	"log"

	"github.com/jlaffaye/ftp"
)

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func handleClientDNSConnection(conn *net.TCPConn, msg messageFormat.MessageFormat) []string {

	dnsResponse := &messageFormat.MessageFormat{}

	encoder := gob.NewEncoder(conn)

	encoder.Encode(msg)

	decoder := gob.NewDecoder(conn)

	decoder.Decode(dnsResponse)

	return dnsResponse.Payload.([]string)
}

func handleClientAuthConnection(conn *net.TCPConn, msgUser messageFormat.MessageFormat) bool {

	encoderServer := gob.NewEncoder(conn)
	encoderServer.Encode(msgUser)

	var authResponse bool

	decoder := gob.NewDecoder(conn)

	decoder.Decode(&authResponse)

	return authResponse
}

func handleServerConnection(reqType string, serverAddr string, fileName string, userName string) {

	data := make([]byte, 1024)
	conn, err := ftp.Dial(serverAddr)
	conn.Login("admin", "admin")
	loginErr := conn.Login("admin", "admin")

	conn.ChangeDir(userName)

	if loginErr != err {
		panic(loginErr)
	}
	if err != nil {
		panic(err)
	}

	if reqType == "str" {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}

		conn.Stor("/"+userName+"/"+fileName, file)
	}

	if reqType == "get" {
		test, _ := conn.Retr(fileName)

		file, _ := os.Create(fileName)
		test.Read(data)
		file.Write(data)
	}

}
func main() {

	// Terminal clear
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println("-------------------------- Codehub --------------------------\n")

	var userAction, login, password, filename string

	fmt.Scanln(&userAction, &login, &password)

	cryptPass := base64Encode([]byte(password))

	if userAction != "auth" {
		fmt.Println("\n")
		log.Fatal("\nAuthentication is needed!\nTry: auth <login> <password>\n")
	}

	tcpAddr, _ := net.ResolveTCPAddr("tcp", "192.168.0.103:2223")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	msg := messageFormat.MessageFormat{
		Origin:  "Client",
		ReqType: "Auth",
	}
	authAddr := handleClientDNSConnection(conn, msg)[0]

	msg = messageFormat.MessageFormat{
		Payload: []string{login, string(cryptPass)},
	}

	tcpAuth, _ := net.ResolveTCPAddr("tcp", authAddr)
	connAuth, _ := net.DialTCP("tcp", nil, tcpAuth)

	authResponse := handleClientAuthConnection(connAuth, msg)

	if !authResponse {
		log.Fatalln("Unauthorized\n")
	}

	msg = messageFormat.MessageFormat{
		Origin:  "Client",
		ReqType: "Server",
	}

	co
	nn, _ = net.DialTCP("tcp", nil, tcpAddr)
	serverAddr := handleClientDNSConnection(conn, msg)
	
	// Client get/str files
	fmt.Println("\nuse get name-of-file.txt to get a code from server")
	fmt.Println("use str name-of-file.txt to upload a code to server\n")
	fmt.Scanln(&userAction/*, &filename*/)

	if userAction != "get" && userAction != "str" {
		fmt.Println("\n")
		log.Fatal("\nWrong command, try again.\n")
	}

	if userAction == "get" {
		handleServerConnection("get", serverAddr[0], filename, login)
	} else {
		handleServerConnection("str", serverAddr[0], filename, login)
	}
}
