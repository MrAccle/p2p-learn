package model

import (
	"bytes"
	"clientBox/utils"
	"github.com/pkg/errors"
	"strings"
)

//cmd(1byte) + LanIP:Port(6byte) + WanIP:Port(6byte)+ NAT type(1byte) + Node type(1byte) + NodeID(32 byte) + UUID(32 byte)
func RegisterProtocal(lanIP string,wanIP string,natType byte,nodeType byte,nodeID string,uuid string)([]byte,error){
	cmd := CMD_REGISTER
	buf := bytes.NewBuffer([]byte{cmd})
	bytNodeID := utils.StringToHex(nodeID)
	bytUUID := utils.StringToHex(uuid)
	if len(bytNodeID) != 32{
		return nil,errors.New("error for node id format")
	}
	if len(bytUUID) != 32{
		return nil,errors.New("error for node id format")
	}
	arrLIP := strings.Split(lanIP,":")
	bytLIP,_ := utils.StrIPToBytes(arrLIP[0])
	bytLPORT,_ := utils.StrPortToBytes(arrLIP[1])

	arrWIP := strings.Split(wanIP,":")
	bytWIP,_ := utils.StrIPToBytes(arrWIP[0])
	bytWPORT,_ := utils.StrPortToBytes(arrWIP[1])

	buf.Write(bytLIP)
	buf.Write(bytLPORT)
	buf.Write(bytWIP)
	buf.Write(bytWPORT)
	buf.Write([]byte{natType})
	buf.Write([]byte{nodeType})
	buf.Write(bytNodeID)
	buf.Write(bytUUID)
	return buf.Bytes(),nil
}


func GetClientProtocal(uuid string,rNodeType byte)([]byte,error){
	var cmd byte
	buf := bytes.NewBuffer([]byte{})
	if rNodeType != NODE_BOX{
		cmd = CMD_GETBOX
	}else{
		cmd = CMD_GETCLIECT
	}
	buf.WriteByte(cmd)
	byt := utils.StringToHex(uuid)
	buf.Write(byt)
	return buf.Bytes(),nil
}