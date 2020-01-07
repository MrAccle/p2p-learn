package model

import (
	"fmt"
	"net"
)

const (
	SERVER_HOST = "0.0.0.0"
	SERVER_PORT = "9111"
)

var server *ServerModel

type ServerModel struct {
	tListener *net.TCPListener
	uListener *net.UDPConn
}

func NewServerModel() *ServerModel {
	tAddr, err := net.ResolveTCPAddr("tcp", SERVER_HOST+":"+SERVER_PORT)
	checkErr(err, "ResolveTCPAddr")
	uAddr, err := net.ResolveUDPAddr("udp", SERVER_HOST+":"+SERVER_PORT)
	checkErr(err, "ResolveUDPAddr")

	tCon, err := net.ListenTCP("tcp", tAddr)
	checkErr(err, "ListenTCP")
	uCon, err := net.ListenUDP("udp", uAddr)
	checkErr(err, "ListenUDP")

	return &ServerModel{
		tListener: tCon,
		uListener: uCon,
	}
}

func init() {
	server = NewServerModel()
}

func (s *ServerModel) HandleUDPConn() {
	buf := make([]byte, 256)
	i, uAddr, err := s.uListener.ReadFromUDP(buf)
	checkErr(err, "listener from udp")
	go handleUDPData(i, uAddr, buf)
}


func handleUDPData(i int, uAddr *net.UDPAddr, data []byte) {
	cmd := data[0]
	switch cmd {
	case CMD_REGISTER:
		handleRegisterCMD(data,uAddr)
	case CMD_GETBOX:
		handleGetBox(data,uAddr)
	case CMD_GETCLIECT:
		handleGetClient(data,uAddr)
	}
}


//Register format
//cmd(1byte) + LanIP:Port(6byte) + WanIP:Port(6byte)+ NAT type(1byte) + Node type(1byte) + NodeID(32 byte) + UUID(32 byte)
func handleRegisterCMD(data []byte,addr *net.UDPAddr){
	bytData := data[1:47]
	bytUUID := data[47:79]
	strUUID := Utils.HexToString(bytUUID)
	//LanIP:Port(6byte) + WanIP:Port(6byte)+ NAT type(1byte) + Node type(1byte) + NodeID(32 byte)
	node,_ := NewNodeWithBytes(bytData)
	if node != nil{
		nodePool.setNewNode(strUUID,node)
	}
	server.uListener.WriteToUDP([]byte("register ack success"),addr)
}

//cmd(1byte)+UUID(32 byte)
func handleGetBox(data []byte,addr *net.UDPAddr){
	bytUUID := data[1:33]
	strUUID := utils.HexToString(bytUUID)
	node,err:= nodePool.getClientNode(strUUID)
	if err != nil{

	}
	fmt.Println(node)
}

//cmd(1byte)+UUID(32 byte)
func handleGetClient(data []byte,addr *net.UDPAddr){
	bytUUID := data[1:33]
	strUUID := utils.HexToString(bytUUID)
	node,err:= nodePool.getBoxNode(strUUID)
	if err != nil{

	}
	fmt.Println(node)
}
func checkErr(err error, tag string) {
	if err != nil {
		fmt.Println(tag, "=>>error in Server model,the reason is:")
		fmt.Println(err.Error())
	}
}
