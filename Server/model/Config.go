package model

//NAT type
const(
	NAT_FULL_CONE  = 0x01+iota
	NAT_RESTRICTED_Cone
	NAT_PORT_RESTRICTED_CONE
	NAT_SYMMETRIC
	NAT_UNKNOW
)

//Node type
const (
	NODE_CLIENT = 0x01 + iota
	NODE_BOX
)

//command type
const (
	CMD_REGISTER byte = 0x01 + iota
	CMD_GETBOX
	CMD_GETCLIECT
)