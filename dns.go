package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"../codehub-sd/messageFormat"
)

type dns struct {
	/*map["Serverx : "1994864:1554""]*/
	table map[string][]string
}

func (d *dns) handleDNSConnection(conn *net.TCPConn) {

	msg := &messageFormat.MessageFormat{}

	decoder := gob.NewDecoder(conn)

	decoder.Decode(msg)

	encoder := gob.NewEncoder(conn)

	defer conn.Close()

	if msg.Origin == "Client" {
		if msg.ReqType == "Auth" {
			msgResponse := &messageFormat.MessageFormat{
				Origin:  "DNS",
				ReqType: "Response",
				Payload: d.table[msg.ReqType],
			}
			encoder.Encode(msgResponse)
		}

		if msg.ReqType == "Server" {

			for _, ipp := range d.table {
				tcpAddrServer, _ := net.ResolveTCPAddr("tcp", ipp[2])
				conn, err := net.DialTCP("tcp", nil, tcpAddrServer)
				encoderServer := gob.NewEncoder(conn)

				if err != nil {
					continue
				}

				msgServer := messageFormat.MessageFormat{Origin: "DNS", ReqType: "ver"}
				response := messageFormat.MessageFormat{Origin: "DNS", ReqType: "Response", Payload: ipp[1]}

				encoderServer.Encode(msgServer)
				fmt.Println("Client requests server address")
				encoder.Encode(response)

				return
			}
		}
	}

	if msg.Origin == "Server" {
		if msg.ReqType == "Hello" {
			d.table[msg.Payload.([]string)[0]] = []string{msg.Payload.([]string)[1], msg.Payload.([]string)[2]}
			fmt.Print("Dns table is: ")
			fmt.Println(d.table)
		}
	}

}

func main() {
	//var dnsTable dns
	dnsTable := &dns{}
	dnsTable.table = make(map[string][]string)
	dnsTable.table["Auth"] = []string{"localhost:1515"}
	fmt.Println("Starting DNS Server...")
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:2223")
	listener, _ := net.ListenTCP("tcp", tcpAddr)

	for {
		tcpConn, _ := listener.AcceptTCP()
		go dnsTable.handleDNSConnection(tcpConn)
	}
}
