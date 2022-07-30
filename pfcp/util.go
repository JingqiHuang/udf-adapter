package pfcp

import (
	"encoding/json"
	"fmt"
	"upf-adapter/config"
	"upf-adapter/pfcp/message"

	"github.com/omec-project/pfcp"
)

func JsonBodyToPfcpAssocReq(body interface{}) pfcp.PFCPAssociationSetupRequest {

	fmt.Println("received PFCPAssociationSetupRequest")
	jsonString, _ := json.Marshal(body)

	s := pfcp.PFCPAssociationSetupRequest{}
	json.Unmarshal(jsonString, &s)

	return s
}

func JsonBodyToPfcpHeartbeatReq(body interface{}) pfcp.HeartbeatRequest {

	fmt.Println("received HeartbeatRequest")
	jsonString, _ := json.Marshal(body)

	s := pfcp.HeartbeatRequest{}
	json.Unmarshal(jsonString, &s)

	return s
}

func JsonBodyToPfcpSessEstReq(body interface{}) pfcp.PFCPSessionEstablishmentRequest {

	fmt.Println("received PFCPSessionEstablishmentRequest")
	jsonString, _ := json.Marshal(body)

	s := pfcp.PFCPSessionEstablishmentRequest{}
	json.Unmarshal(jsonString, &s)

	return s
}

func JsonBodyToPfcpSessModReq(body interface{}) pfcp.PFCPSessionModificationRequest {

	fmt.Println("received PFCPSessionModificationRequest")
	jsonString, _ := json.Marshal(body)

	s := pfcp.PFCPSessionModificationRequest{}
	json.Unmarshal(jsonString, &s)

	return s
}

func JsonBodyToPfcpSessDelReq(body interface{}) pfcp.PFCPSessionDeletionRequest {

	fmt.Println("received PFCPSessionDeletionRequest")
	jsonString, _ := json.Marshal(body)

	s := pfcp.PFCPSessionDeletionRequest{}
	json.Unmarshal(jsonString, &s)

	return s
}

func ForwardPfcpMsgToUpf(udpPodMsg config.UdpPodPfcpMsg) ([]byte, error) {

	pMsg := udpPodMsg.Msg
	nodeId := udpPodMsg.UpNodeID
	pfcpTxnChan := make(config.PfcpTxnChan)
	//identify msg type
	switch pMsg.Header.MessageType {
	case pfcp.PFCP_ASSOCIATION_SETUP_REQUEST:
		s := JsonBodyToPfcpAssocReq(pMsg.Body)

		//store txn in seq:chan map
		config.InsertUpfPfcpTxn(pMsg.Header.SequenceNumber, pfcpTxnChan)
		pMsg.Body = s
		message.SendPfcpAssociationSetupRequest(nodeId, pMsg)

	case pfcp.PFCP_HEARTBEAT_REQUEST:
		s := JsonBodyToPfcpHeartbeatReq(pMsg.Body)

		//store txn in seq:chan map
		config.InsertUpfPfcpTxn(pMsg.Header.SequenceNumber, pfcpTxnChan)
		pMsg.Body = s
		message.SendHeartbeatRequest(nodeId, pMsg)
	case pfcp.PFCP_SESSION_ESTABLISHMENT_REQUEST:
		s := JsonBodyToPfcpSessEstReq(pMsg.Body)

		//store txn in seq:chan map
		config.InsertUpfPfcpTxn(pMsg.Header.SequenceNumber, pfcpTxnChan)
		pMsg.Body = s
		message.SendPfcpSessionEstablishmentRequest(nodeId, pMsg)
	case pfcp.PFCP_SESSION_MODIFICATION_REQUEST:
		s := JsonBodyToPfcpSessModReq(pMsg.Body)

		//store txn in seq:chan map
		config.InsertUpfPfcpTxn(pMsg.Header.SequenceNumber, pfcpTxnChan)
		pMsg.Body = s
		message.SendPfcpSessionModificationRequest(nodeId, pMsg)
	case pfcp.PFCP_SESSION_DELETION_REQUEST:
		s := JsonBodyToPfcpSessDelReq(pMsg.Body)

		//store txn in seq:chan map
		config.InsertUpfPfcpTxn(pMsg.Header.SequenceNumber, pfcpTxnChan)
		pMsg.Body = s
		message.SendPfcpSessionDeletionRequest(nodeId, pMsg)
	default:
		return nil, fmt.Errorf("invalid msg type [%v] from smf", pMsg.Header.MessageType)
	}

	//wait for response from UPF
	pfcpRsp := <-pfcpTxnChan

	return pfcpRsp, nil
}
