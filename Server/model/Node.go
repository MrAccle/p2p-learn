package model

import (
	"github.com/pkg/errors"
	"server/utils"
)


//one connection point is a Node
type Node struct {
	//node local area network ip address and ip port
	LanIP   string
	LanPort int32
	//node wide area network ip address and ip port
	WanIP string
	WanPort int32
	//the node network nat type
	//can be set
	NatType byte
	//the node type
	NodeType byte
	//node unique id
	NodeID string
}

func NewNode(lanIP string, lanPort int32, wanIP string, wanPort int32, natType byte, nodeType byte, nodeID string) *Node {
	return &Node{LanIP: lanIP, LanPort: lanPort, WanIP: wanIP, WanPort: wanPort, NatType: natType, NodeType: nodeType, NodeID: nodeID}
}


//this function is use a byte array to create a new node
//the byte format like this
//LanIP:Port(6byte) + WanIP:Port(6byte)+ NAT type(1byte) + Node type(1byte) + NodeID(32 byte)
func NewNodeWithBytes(data []byte)(*Node,error){
	if len(data)!=46{
		return nil,errors.New("error for data")
	}
	byteLanIP := data[0:4]
	byteLanPort := data[4:6]
	byteWanIP := data[6:10]
	byteWanPort := data[10:12]
	byteNatType := data[12:13]
	byteNodeType := data[13:14]
	byteNodeID := data[14:]

	strLanIP,_ := utils.IPBytesToString(byteLanIP)
	strLanPort := utils.PortBytesToInt32(byteLanPort)

	strWanIP,_ := utils.IPBytesToString(byteWanIP)
	strWanPort := utils.PortBytesToInt32(byteWanPort)

	strNodeID := utils.HexToString(byteNodeID)

	return NewNode(strLanIP,strLanPort,strWanIP,strWanPort,byteNatType[0],byteNodeType[0],strNodeID),nil
}
