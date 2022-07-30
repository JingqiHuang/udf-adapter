package handler

import (
	"encoding/json"
	"fmt"
	"upf-adapter/config"

	"github.com/omec-project/pfcp"
	"github.com/omec-project/pfcp/pfcpUdp"
)

func encodeAndSendRsp(msg *pfcpUdp.Message) error {
	var pMsgBody interface{}
	switch msg.PfcpMessage.Header.MessageType {
	case pfcp.PFCP_ASSOCIATION_SETUP_RESPONSE:
		pMsgBody = msg.PfcpMessage.Body.(pfcp.PFCPAssociationSetupResponse)
	case pfcp.PFCP_HEARTBEAT_RESPONSE:
		pMsgBody = msg.PfcpMessage.Body.(pfcp.HeartbeatResponse)
	case pfcp.PFCP_SESSION_ESTABLISHMENT_RESPONSE:
		pMsgBody = msg.PfcpMessage.Body.(pfcp.PFCPSessionEstablishmentResponse)
		fmt.Println("encodeAndSendRsp PFCPSessionEstablishmentResponse")
	case pfcp.PFCP_SESSION_MODIFICATION_RESPONSE:
		pMsgBody = msg.PfcpMessage.Body.(pfcp.PFCPSessionModificationResponse)
	case pfcp.PFCP_SESSION_DELETION_RESPONSE:
		pMsgBody = msg.PfcpMessage.Body.(pfcp.PFCPSessionDeletionResponse)
	default:
		return fmt.Errorf("invalid rsp msg type %v", msg.PfcpMessage.Header.MessageType)
	}

	msg.PfcpMessage.Body = pMsgBody

	pRspJson, err := json.Marshal(msg.PfcpMessage)
	if err != nil {
		fmt.Printf("\njson encode error: %v", err)
	}

	//Get the PFCP Txn
	pfcpTxnChan := config.GetUpfPfcpTxn(msg.PfcpMessage.Header.SequenceNumber)

	//Send Rsp back to http txn
	pfcpTxnChan <- pRspJson

	return nil
}

func HandlePfcpAssociationSetupResponse(msg *pfcpUdp.Message) {
	//Encode pfcp rsp to json and send to http txn
	if err := encodeAndSendRsp(msg); err != nil {
		fmt.Printf("handle pfcp association response error : %v", err)
	}
}

func HandlePfcpHeartbeatResponse(msg *pfcpUdp.Message) {
	//Encode pfcp rsp to json and send to http txn
	if err := encodeAndSendRsp(msg); err != nil {
		fmt.Printf("handle pfcp heartbeat response error : %v", err)
	}
}

func HandlePfcpSessionEstablishmentResponse(msg *pfcpUdp.Message) {
	//Encode pfcp rsp to json and send to http txn
	if err := encodeAndSendRsp(msg); err != nil {
		fmt.Printf("handle pfcp session establishment response error : %v", err)
	}
}

func HandlePfcpSessionModificationResponse(msg *pfcpUdp.Message) {
	//Encode pfcp rsp to json and send to http txn
	if err := encodeAndSendRsp(msg); err != nil {
		fmt.Printf("handle pfcp session modify response error : %v", err)
	}
}

func HandlePfcpSessionDeletionResponse(msg *pfcpUdp.Message) {
	//Encode pfcp rsp to json and send to http txn
	if err := encodeAndSendRsp(msg); err != nil {
		fmt.Printf("handle pfcp session delete response error : %v", err)
	}
}
