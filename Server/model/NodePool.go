package model

import (
	"container/list"
	"github.com/pkg/errors"
	"sync"
)

/*
	this model version is 1.1 bate
	and we acquiesce the user only have one Client and one box
*/
var(
	nodePool *NodePool
)
type NodePool struct {
	pool map[string]*list.List
	sync.RWMutex
}

func init(){
	nodePool = NewNodePool()
}

func NewNodePool() *NodePool {
	return &NodePool{}
}

func (n *NodePool)setNewNode(uuid string,node *Node){
	n.RLock()
	k,ok := n.pool[uuid]
	n.RUnlock()
	if !ok{
		k = list.New()
	}
	n.Lock()
	k.PushBack(node)
	n.Unlock()
}

//this function only get the first box node
func (n *NodePool)getBoxNode(uuid string)(*Node,error){
	n.RLock()
	k,ok := n.pool[uuid]
	n.RUnlock()
	if !ok{
		return nil,errors.New("did't have this uuid node")
	}
	var temp *Node
	for e := k.Front();e != nil ;e.Next(){
		if e.Value.(*Node).NodeType == NODE_BOX {
			temp = e.Value.(*Node)
			break
		}
	}
	if temp != nil{
		return temp,nil
	}else{
		return nil,errors.New("did't have box node")
	}
}

//this function only get the first Client node
func (n *NodePool)getClientNode(uuid string)(*Node,error){
	n.RLock()
	k,ok := n.pool[uuid]
	n.RUnlock()
	if !ok{
		return nil,errors.New("did't have this uuid node")
	}
	var temp *Node
	for e := k.Front();e != nil ;e.Next(){
		if e.Value.(*Node).NodeType == NODE_CLIENT {
			temp = e.Value.(*Node)
			break
		}
	}
	if temp != nil{
		return temp,nil
	}else{
		return nil,errors.New("did't have box node")
	}
}
