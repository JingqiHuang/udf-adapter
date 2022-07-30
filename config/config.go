package config

import (
	"fmt"
	"net"
	"sync"
	"time"

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
