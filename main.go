package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"upf-adapter/config"
	"upf-adapter/pfcp/udp"

	"upf-adapter/pfcp"
)

// var seq uint32

// func getSeqNumber() uint32 {
// 	return atomic.AddUint32(&seq, 1)
// }

func rec_msg(w http.ResponseWriter, req *http.Request) {

	for headerName, headerValue := range req.Header {
		fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
	}

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("server: could not read request body: %s\n", err)
	}

	// reconstruct pfcp msg
	//fmt.Printf("server: request body: %s\n", reqBody)

	var udpPodMsg config.UdpPodPfcpMsg
	json.Unmarshal(reqBody, &udpPodMsg)

	fmt.Printf("pfcp MessageType %v", udpPodMsg.Msg.Header.MessageType)
	//fmt.Printf("s type %T\n", s)

	fmt.Printf("s.NodeID type %s\n", udpPodMsg.UpNodeID.NodeIdValue)
	//fmt.Printf("s.RecoveryTimeStamp type %T\n", s.RecoveryTimeStamp)
	//fmt.Printf("s.UPFunctionFeatures type %T\n", s.UPFunctionFeatures)
	//fmt.Printf("s.CPFunctionFeatures type %T\n", s.CPFunctionFeatures)
	//fmt.Printf("s.UserPlaneIPResourceInformation type %T\n", s.UserPlaneIPResourceInformation)

	//fmt.Printf("s val %v\n", s)

	//fmt.Printf("s.NodeID val %v\n", s.NodeID)
	//fmt.Printf("s.RecoveryTimeStamp val %v\n", s.RecoveryTimeStamp)
	//fmt.Printf("s.UPFunctionFeatures val %v\n", s.UPFunctionFeatures)
	//fmt.Printf("s.CPFunctionFeatures val %v\n", s.CPFunctionFeatures)
	//fmt.Printf("s.UserPlaneIPResourceInformation val %v\n", s.UserPlaneIPResourceInformation)

	pfcpJsonRsp, err := pfcp.ForwardPfcpMsgToUpf(udpPodMsg)
	if err != nil {
		fmt.Println("Error HttpLib received pfcp Rsp ")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(pfcpJsonRsp)
	fmt.Println("Response Sent for ", udpPodMsg.Msg.Header.MessageType)
}

func init() {
	udp.Run(pfcp.Dispatch)
}

func main() {
	http.HandleFunc("/", rec_msg)
	http.ListenAndServe(":8090", nil)

}
