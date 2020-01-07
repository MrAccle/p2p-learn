package model

import (

	"fmt"
	"net"
	reuse "github.com/libp2p/go-reuseport"
)



//tcp connect is to register node
//udp connect is to though nat
type serverConnect struct {
	TCon net.Conn
	UCon net.Conn
	TLocalAddr net.Addr
}

func MewServerConnect() *serverConnect {
	tcon,err := reuse.Dial("tcp",LOCAL_HOST+":"+LOCAL_TCP_PORT,SERVER_HOST+":"+SERVER_PORT)
	checkError(err,"reuse dial tcp")
	ucon,err := reuse.Dial("udp",LOCAL_HOST+":"+LOCAL_UDP_PORT,SERVER_HOST+":"+SERVER_PORT)
	checkError(err,"reuse dial udp")
	//tcon,err := net.DialTCP()

	return &serverConnect{
		TCon:tcon,
		UCon:ucon,
	}
}

func (s* serverConnect)Register(){
	//localNet := ""

	buf := RegisterProtocal()
	s.TCon.Write(buf)
}
func checkError(err error,tag string){
	if err != nil{
		fmt.Println(tag,err.Error())
	}
}