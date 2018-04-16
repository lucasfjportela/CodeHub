package main

import (
	"codehub-sd/messageFormat"
	"encoding/gob"
	"net"
	"os"

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

	if reqType == "stor" {
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

	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:2223")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	msg := messageFormat.MessageFormat{
		Origin:  "Client",
		ReqType: "Auth",
	}
	authAddr := handleClientDNSConnection(conn, msg)

	msg = messageFormat.MessageFormat{
		Origin:  "Client",
		ReqType: "Auth",
		Payload: []string{"Denini", "123"},
	}

	tcpAuth, _ := net.ResolveTCPAddr("tcp", authAddr)
	connAuth, _ := net.DialTCP("tcp", nil, tcpAuth)

	authResponse := handleClientAuthConnection(connAuth, msg)

	if !authResponse {
		panic("You can't auth")
	}

	msg = messageFormat.MessageFormat{
		Origin:  "Client",
		ReqType: "Server",
	}

	serverAddr := handleClientDNSConnection(conn, msg)

	handleServerConnection("get", serverAddr)

}
