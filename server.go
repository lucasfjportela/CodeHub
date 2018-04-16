package main

import (
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
			if msg.ReqType == "sto" {
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

	var tcpConn net.TCPConn

	defer tcpConn.Close()

	for {
		fmt.Println("Server listening dns on port 1111...")
		listener.AcceptTCP()
		fmt.Println("ROLA")
	}
}

func main() {
	/*fmt.Println("Starting Server...")

	 */
	factory := &filedriver.FileDriverFactory{
		RootPath: "C:/Users/Matheus/Desktop/ROLA",
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
