package udp

import (
	"fmt"
	"net"
	"time"

	"github.com/omec-project/pfcp"
	"github.com/omec-project/pfcp/pfcpType"
	"github.com/omec-project/pfcp/pfcpUdp"
)

var Server *pfcpUdp.PfcpServer

var ServerStartTime time.Time
var CPNodeID *pfcpType.NodeID

func init() {
	CPNodeID = &pfcpType.NodeID{NodeIdType: uint8(2), NodeIdValue: []byte("upf-adapter")}
}

func SendPfcp(msg pfcp.Message, addr *net.UDPAddr, eventData interface{}) error {
	err := Server.WriteTo(msg, addr, eventData)
	if err != nil {
		fmt.Printf("Failed to send PFCP message: %v", err)

		return err
	}

	return nil
}

func Run(Dispatch func(*pfcpUdp.Message)) {

	Server = pfcpUdp.NewPfcpServer(CPNodeID.ResolveNodeIdToIp().String())

	err := Server.Listen()
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
	}
	fmt.Printf("Listen on %s", Server.Conn.LocalAddr().String())

	go func(p *pfcpUdp.PfcpServer) {
		for {
			var pfcpMessage pfcp.Message
			remoteAddr, eventData, err := p.ReadFrom(&pfcpMessage)
			if err != nil {
				if err.Error() == "Receive resend PFCP request" {
					fmt.Println(err)
				} else {
					fmt.Printf("Read PFCP error: %v", err)
				}

				continue
			}

			msg := pfcpUdp.NewMessage(remoteAddr, &pfcpMessage, eventData)
			go Dispatch(&msg)
		}
	}(Server)

	ServerStartTime = time.Now()
}
