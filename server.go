package main

import (
	"codehub-sd/messageFormat"
	"encoding/gob"
	"fmt"
	"net"

	filedriver "github.com/goftp/file-driver"
	"github.com/goftp/server"
)

type TestAuth struct {
	Name     string
	Password string
}

func (a *TestAuth) CheckPasswd(name, pass string) (bool, error) {
	if name != a.Name || pass != a.Password {
		return false, nil
	}
	return true, nil
}

/*Func that execute the server operation*/
func handleServerConnection(conn *net.TCPConn) {

	/*var data map[string]string

	decoder := gob.NewDecoder(conn)

	decoder.Decode(&data)

	//readed := make([]byte, 1024)
	//datasize, _ := conn.Read(readed)
	//data := readed[:datasize]
	fmt.Println(data)
	*/
	/*
		msg := messageFormat.MessageFormat{}

		decoder := gob.NewDecoder(conn)
		decoder.Decode(&msg)

		if msg.Origin == "Client" {
			if msg.ReqType == "str" {
				fmt.Println("Aqui guarda o arquivo")
			}

			if msg.ReqType == "get" {
				fmt.Println("Aqui dá o arquivo pra ele")
			}
		}*/

}

func HandleServerDNSConnection() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:1111")
	listener, _ := net.ListenTCP("tcp", tcpAddr)

	for {
		fmt.Println("Server listening dns on port 1111...")
		listener.AcceptTCP()
		fmt.Println()
	}
}

func ServerHello() {

	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:2424")
	tcpDNS, _ := net.ResolveTCPAddr("tcp", "localhost:2223")
	conn, _ := net.DialTCP("tcp", tcpAddr, tcpDNS)

	encoder := gob.NewEncoder(conn)
	msg := &messageFormat.MessageFormat{
		Origin:  "Server",
		ReqType: "Hello",
		Payload: []string{"Server1", "localhost:2121", "localhost:1111"},
	}

	encoder.Encode(msg)

	fmt.Println("Sending hello to dns server \\o")

}

func main() {

	ServerHello()

	factory := &filedriver.FileDriverFactory{
		RootPath: "//10.1.4.65/Users/Public/codehub",
		Perm:     server.NewSimplePerm("root", "root"),
	}

	opts := &server.ServerOpts{
		Factory:  factory,
		Port:     2121,
		Hostname: "localhost",

		Auth: &TestAuth{
			Name:     "admin",
			Password: "admin",
		},
	}

	go HandleServerDNSConnection()

	serv := server.NewServer(opts)

	serv.ListenAndServe()
}
