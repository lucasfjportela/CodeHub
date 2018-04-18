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

func HandleServerDNSConnection() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "192.168.0.110:1111")
	listener, _ := net.ListenTCP("tcp", tcpAddr)

	for {
		fmt.Println("Server 1 listening dns on port 1111...")
		listener.AcceptTCP()
		fmt.Println()
	}
}

func ServerHello() {

	tcpAddr, _ := net.ResolveTCPAddr("tcp", "192.168.0.110:2424")
	tcpDNS, _ := net.ResolveTCPAddr("tcp", "192.168.0.103:2223")
	conn, _ := net.DialTCP("tcp", tcpAddr, tcpDNS)

	encoder := gob.NewEncoder(conn)
	msg := &messageFormat.MessageFormat{
		Origin:  "Server",
		ReqType: "Hello",
		Payload: []string{"Server1", "192.168.0.110:2121", "192.168.0.110:1111"},
	}

	encoder.Encode(msg)

	fmt.Println("Sending hello to dns server \\o")

}

func main() {

	ServerHello()

	factory := &filedriver.FileDriverFactory{
		RootPath: "//192.168.0.100/C:/Users/Public/Codehub",
		Perm:     server.NewSimplePerm("root", "root"),
	}

	opts := &server.ServerOpts{
		Factory:  factory,
		Port:     2121,
		Hostname: "192.168.0.110",

		Auth: &TestAuth{
			Name:     "admin",
			Password: "admin",
		},
	}

	go HandleServerDNSConnection()

	serv := server.NewServer(opts)

	serv.ListenAndServe()
}