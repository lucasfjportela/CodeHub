package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"../codehub-sd/messageFormat"
)

type dns struct {
	/*map["Serverx : "1994864:1554""]*/
	table map[string]string
}

func (d *dns) handleDNSConnection(conn *net.TCPConn) {

	msg := &messageFormat.MessageFormat{}

	decoder := gob.NewDecoder(conn)

	decoder.Decode(msg)

	encoder := gob.NewEncoder(conn)

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
				tcpAddrServer, _ := net.ResolveTCPAddr("tcp", ipp)
				conn, err := net.DialTCP("tcp", nil, tcpAddrServer)
				encoderServer := gob.NewEncoder(conn)

				if err != nil {
					continue
				}

				msgServer := messageFormat.MessageFormat{Origin: "DNS", ReqType: "ver"}
				response := messageFormat.MessageFormat{Origin: "DNS", ReqType: "Response", Payload: ipp}

				encoderServer.Encode(msgServer)
				fmt.Println("Client requests server address")
				encoder.Encode(response)

				conn.Close()
				return
			}
		}
	}

}

func main() {
	//var dnsTable dns
	dnsTable := &dns{}
	dnsTable.table = make(map[string]string)
	dnsTable.table["Server1"] = "localhost:1111"
	dnsTable.table["Server2"] = "localhost:1112"
	dnsTable.table["Server3"] = "localhost:1113"
	dnsTable.table["Auth"] = "localhost:1515"

	fmt.Println("Starting DNS Server...")
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:2223")
	listener, _ := net.ListenTCP("tcp", tcpAddr)

	var tcpConn net.TCPConn
	defer tcpConn.Close()

	for {
		tcpConn, _ := listener.AcceptTCP()
		go dnsTable.handleDNSConnection(tcpConn)
	}
}
