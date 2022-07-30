package message

import (
	"net"

	"upf-adapter/pfcp/udp"

	"github.com/omec-project/pfcp"
	"github.com/omec-project/pfcp/pfcpType"
	"github.com/omec-project/pfcp/pfcpUdp"
)

const MaxPfcpUdpDataSize = 1024

const (
	PFCP_PORT        = 8805
	PFCP_MAX_UDP_LEN = 2048
)

func SendPfcpAssociationSetupRequest(upNodeID pfcpType.NodeID, pMsg pfcp.Message) error {

	message := pfcp.Message{
		Header: pfcp.Header{
			Version:        pfcp.PfcpVersion,
			MP:             0,
			S:              pfcp.SEID_NOT_PRESENT,
			MessageType:    pfcp.PFCP_ASSOCIATION_SETUP_REQUEST,
			SequenceNumber: pMsg.Header.SequenceNumber,
		},
		Body: pMsg.Body,
	}

	addr := &net.UDPAddr{
		IP:   upNodeID.ResolveNodeIdToIp(),
		Port: PFCP_PORT,
	}

	udp.SendPfcp(message, addr, nil)
	return nil
}

func SendHeartbeatRequest(upNodeID pfcpType.NodeID, pMsg pfcp.Message) error {
	message := pfcp.Message{
		Header: pfcp.Header{
			Version:        pfcp.PfcpVersion,
			MP:             0,
			S:              pfcp.SEID_NOT_PRESENT,
			MessageType:    pfcp.PFCP_HEARTBEAT_REQUEST,
			SequenceNumber: pMsg.Header.SequenceNumber,
		},
		Body: pMsg.Body,
	}

	addr := &net.UDPAddr{
		IP:   upNodeID.ResolveNodeIdToIp(),
		Port: pfcpUdp.PFCP_PORT,
	}

	udp.SendPfcp(message, addr, nil)
	return nil
}

func SendPfcpSessionEstablishmentRequest(upNodeID pfcpType.NodeID, pMsg pfcp.Message) error {
	ip := upNodeID.ResolveNodeIdToIp()

	message := pfcp.Message{
		Header: pfcp.Header{
			Version:         pfcp.PfcpVersion,
			MP:              1,
			S:               pfcp.SEID_PRESENT,
			MessageType:     pfcp.PFCP_SESSION_ESTABLISHMENT_REQUEST,
			SEID:            0,
			SequenceNumber:  pMsg.Header.SequenceNumber,
			MessagePriority: 0,
		},
		Body: pMsg.Body,
	}

	upaddr := &net.UDPAddr{
		IP:   ip,
		Port: pfcpUdp.PFCP_PORT,
	}

	udp.SendPfcp(message, upaddr, nil)
	return nil
}

func SendPfcpSessionModificationRequest(upNodeID pfcpType.NodeID, pMsg pfcp.Message) error {
	ip := upNodeID.ResolveNodeIdToIp()
	remoteSEID := pMsg.Header.SEID
	message := pfcp.Message{
		Header: pfcp.Header{
			Version:         pfcp.PfcpVersion,
			MP:              1,
			S:               pfcp.SEID_PRESENT,
			MessageType:     pfcp.PFCP_SESSION_MODIFICATION_REQUEST,
			SEID:            remoteSEID,
			SequenceNumber:  pMsg.Header.SequenceNumber,
			MessagePriority: 12,
		},
		Body: pMsg.Body,
	}

	upaddr := &net.UDPAddr{
		IP:   ip,
		Port: pfcpUdp.PFCP_PORT,
	}

	udp.SendPfcp(message, upaddr, nil)
	return nil
}

func SendPfcpSessionDeletionRequest(upNodeID pfcpType.NodeID, pMsg pfcp.Message) error {
	ip := upNodeID.ResolveNodeIdToIp()
	remoteSEID := pMsg.Header.SEID
	message := pfcp.Message{
		Header: pfcp.Header{
			Version:         pfcp.PfcpVersion,
			MP:              1,
			S:               pfcp.SEID_PRESENT,
			MessageType:     pfcp.PFCP_SESSION_DELETION_REQUEST,
			SEID:            remoteSEID,
			SequenceNumber:  pMsg.Header.SequenceNumber,
			MessagePriority: 12,
		},
		Body: pMsg.Body,
	}

	upaddr := &net.UDPAddr{
		IP:   ip,
		Port: pfcpUdp.PFCP_PORT,
	}

	udp.SendPfcp(message, upaddr, nil)
	return nil
}
