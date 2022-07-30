package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"upf-adapter/config"
	"upf-adapter/pfcp/udp"

	"upf-adapter/pfcp"

	omec_pfcp "github.com/omec-project/pfcp"
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

	pfcpJsonRsp, err := pfcp.ForwardPfcpMsgToUpf(udpPodMsg)

	// []byte -> struct
	pfcpRspMsg := omec_pfcp.Message{}
	json.Unmarshal(pfcpJsonRsp, &pfcpRspMsg)

	//- struct -> []byte for customized struct
	UdpRsp := &config.UdpPodPfcpRspMsg{
		Msg: pfcpRspMsg,
	}

	if pfcpRspMsg.Header.MessageType == omec_pfcp.PFCP_SESSION_ESTABLISHMENT_RESPONSE {
		fmt.Printf("\ngot PFCP_SESSION_ESTABLISHMENT_RESPONSE\n")

		d := json.NewDecoder(bytes.NewBuffer(pfcpJsonRsp))
		d.UseNumber()

		dat := make(map[string]interface{})
		if err := d.Decode(&dat); err != nil {
			panic(err)
		}
		fmt.Printf("\ndat %v", dat)

		bodys := dat["Body"].(map[string]interface{})
		fmt.Printf("\n bodys %v", bodys)

		upFseid := bodys["UPFSEID"].(map[string]interface{})

		seidJsonNum := upFseid["Seid"].(json.Number)
		seid64, _ := strconv.ParseUint(string(seidJsonNum), 10, 64)
		seidStr := config.SeidConv(seid64)
		UdpRsp.CustomMsg.Fseid = seidStr

	}

	udpRspJsonByte, _ := json.Marshal(UdpRsp)
	if err != nil {
		fmt.Println("Error HttpLib received pfcp Rsp ")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(udpRspJsonByte)

	fmt.Println("Response Sent for ", udpPodMsg.Msg.Header.MessageType)
}

func init() {
	udp.Run(pfcp.Dispatch)
}

func main() {
	http.HandleFunc("/", rec_msg)
	http.ListenAndServe(":8090", nil)

}
