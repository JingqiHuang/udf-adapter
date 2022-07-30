package config

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"

	// "upf-adapter/pfcp/util"

	"github.com/omec-project/pfcp"
	"github.com/omec-project/pfcp/pfcpType"
)

type UPFStatus int

const MaxUpfProbeRetryInterval time.Duration = 5 //Seconds

var UpfCfg Config

const (
	NotAssociated          UPFStatus = 0
	AssociatedSettingUp    UPFStatus = 1
	AssociatedSetUpSuccess UPFStatus = 2
)

type UPNode struct {
	UpfName string
	NodeID  pfcpType.NodeID
	ANIP    net.IP
	State   UPFStatus
}

type Config struct {
	UpfLock sync.RWMutex
	UPFs    map[string]*UPNode
}

type UdpPodMsgType int

type UdpPodPfcpMsg struct {
	SEID     string
	SmfIp    string
	UpNodeID pfcpType.NodeID
	// message type contains in Msg.Header
	Msg       pfcp.Message
	Addr      *net.UDPAddr
	EventData interface{}
}

type CustomFieldsUdpPodPfcpRspMsg struct {
	Fseid string `json:"fseid"`
}

type UdpPodPfcpRspMsg struct {
	// message type contains in Msg.Header
	Msg       pfcp.Message                 `json:"msg"`
	CustomMsg CustomFieldsUdpPodPfcpRspMsg `json:"-"`
}

func SeidConv(seid uint64) (seidStr string) {
	seidStr = strconv.FormatUint(seid, 16)
	return seidStr
}

func (udpPodPfcpRspMsg *UdpPodPfcpRspMsg) MarshalJSON() ([]byte, error) {
	type Alias UdpPodPfcpRspMsg

	fseidStr := "0"

	if udpPodPfcpRspMsg.Msg.Header.MessageType == pfcp.PFCP_SESSION_ESTABLISHMENT_RESPONSE {
		fseidStr = udpPodPfcpRspMsg.CustomMsg.Fseid
	}

	customMsg := CustomFieldsUdpPodPfcpRspMsg{
		Fseid: fseidStr,
	}

	return json.Marshal(&struct {
		CustomMsg CustomFieldsUdpPodPfcpRspMsg `json:"customFieldsUdpPodPfcpMsg"`
		*Alias
	}{
		CustomMsg: customMsg,
		Alias:     (*Alias)(udpPodPfcpRspMsg),
	})
}

type PfcpTxnChan chan []byte

var UpfTxns map[uint32]PfcpTxnChan

func init() {
	UpfCfg = Config{
		UPFs: make(map[string]*UPNode),
	}

	UpfTxns = make(map[uint32]PfcpTxnChan)

}

func InsertUpfNode(upfName string) {
	UpfCfg.UpfLock.Lock()
	defer UpfCfg.UpfLock.Unlock()

	//if UPF is already not added
	if _, ok := UpfCfg.UPFs[upfName]; !ok {

		upf := UPNode{
			UpfName: upfName,
			State:   NotAssociated,
			//IP/NodeId : Todo
		}
		UpfCfg.UPFs[upfName] = &upf
	}

}

func RemoveUpfNode(upfName string) {
	UpfCfg.UpfLock.Lock()
	defer UpfCfg.UpfLock.Unlock()

	if upf, ok := UpfCfg.UPFs[upfName]; ok {
		delete(UpfCfg.UPFs, upf.UpfName)
	}
}

func InsertUpfPfcpTxn(seq uint32, pfcpTxnChan PfcpTxnChan) {
	fmt.Printf("\n InsertUpfPfcpTxn Seq[%v]", seq)
	UpfTxns[seq] = pfcpTxnChan
}

func GetUpfPfcpTxn(seq uint32) PfcpTxnChan {
	pfcpTxnChan := UpfTxns[seq]
	if pfcpTxnChan != nil {
		delete(UpfTxns, seq)
		fmt.Printf("\n GetUpfPfcpTxn Seq[%v] found", seq)
		return pfcpTxnChan
	}
	fmt.Printf("\n GetUpfPfcpTxn Seq[%v] not found", seq)
	return nil
}

type CustomFieldsUdpPodPfcpMsg struct {
	Msg   pfcp.Message `json:"msg"`
	Fseid string       `json:"fseid"`
}

func (udpPodPfcpMsg *UdpPodPfcpMsg) UnmarshalJSON(data []byte) error {
	type Alias UdpPodPfcpMsg
	auxCustom := &struct {
		CustomMsg CustomFieldsUdpPodPfcpMsg `json:"customFieldsUdpPodPfcpMsg"`
		*Alias
	}{
		Alias: (*Alias)(udpPodPfcpMsg),
	}
	if err := json.Unmarshal(data, &auxCustom); err != nil {
		fmt.Printf("UdpPodPfcpMsg unmarshall error %v\n", err)
		return err
	}
	fseidStr := auxCustom.CustomMsg.Fseid

	fseid, _ := strconv.ParseUint(fseidStr, 16, 64)

	udpPodPfcpMsg.Msg = auxCustom.CustomMsg.Msg

	if udpPodPfcpMsg.Msg.Header.MessageType == pfcp.PFCP_SESSION_ESTABLISHMENT_REQUEST {

		jsonString, _ := json.Marshal(udpPodPfcpMsg.Msg.Body)

		body := pfcp.PFCPSessionEstablishmentRequest{}
		json.Unmarshal(jsonString, &body)

		body.CPFSEID.Seid = fseid
		udpPodPfcpMsg.Msg.Body = body
	} else if udpPodPfcpMsg.Msg.Header.MessageType == pfcp.PFCP_SESSION_DELETION_REQUEST {

		jsonString, _ := json.Marshal(udpPodPfcpMsg.Msg.Body)
		body := pfcp.PFCPSessionDeletionRequest{}
		json.Unmarshal(jsonString, &body)

		body.CPFSEID.Seid = fseid
		udpPodPfcpMsg.Msg.Body = body
	} else if udpPodPfcpMsg.Msg.Header.MessageType == pfcp.PFCP_SESSION_MODIFICATION_REQUEST {
		jsonString, _ := json.Marshal(udpPodPfcpMsg.Msg.Body)
		body := pfcp.PFCPSessionModificationRequest{}
		json.Unmarshal(jsonString, &body)

		body.CPFSEID.Seid = fseid
		udpPodPfcpMsg.Msg.Body = body
	}
	return nil
}
