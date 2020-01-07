package Utils

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

//0低位 1高位
func IntToUint16(len int) [2]byte {
	buf := [2]byte{0x00, 0x00}
	for i := 0; i < 8; i++ {
		temp := 0x01 << uint(i)
		if (len & temp) == temp {
			buf[0] = buf[0] | (0x01 << uint(i))
		}
	}
	//高位
	for i := 0; i < 8; i++ {
		temp := 0x01 << uint(i)
		if (len & temp) == temp {
			buf[1] = buf[1] | (0x01 << uint(i))
		}
	}
	return buf
}

func Uint16ToInt(len [2]byte)int{
	var data int
	data |= int(len[1]<<8)
	data |= int(len[0])
	return data
}
//uuid or node
func StringToHex(data string) []byte {
	//tmep := []byte{}
	buf := bytes.NewBuffer([]byte{})
	bytBuf := []rune(data)
	len := len(bytBuf)
	if len%2 != 0 {
		return nil
	}
	//fmt.Println(buf.Len())
	for i := 0; i < len; i += 2 {
		//fmt.Println(i)
		var runeA rune
		var runeB rune
		switch bytBuf[i] {
		case '0':
			runeA = 0x00
		case '1':
			runeA = 0x10
		case '2':
			runeA = 0x20
		case '3':
			runeA = 0x30
		case '4':
			runeA = 0x40
		case '5':
			runeA = 0x50
		case '6':
			runeA = 0x60
		case '7':
			runeA = 0x70
		case '8':
			runeA = 0x80
		case '9':
			runeA = 0x90
		case 'A', 'a':
			runeA = 0xA0
		case 'B', 'b':
			runeA = 0xB0
		case 'C', 'c':
			runeA = 0xC0
		case 'D', 'd':
			runeA = 0xD0
		case 'E', 'e':
			runeA = 0xE0
		case 'F', 'f':
			runeA = 0xF0
		}
		switch bytBuf[i+1] {
		case '0':
			runeB = 0x00
		case '1':
			runeB = 0x01
		case '2':
			runeB = 0x02
		case '3':
			runeB = 0x03
		case '4':
			runeB = 0x04
		case '5':
			runeB = 0x05
		case '6':
			runeB = 0x06
		case '7':
			runeB = 0x07
		case '8':
			runeB = 0x08
		case '9':
			runeB = 0x09
		case 'A', 'a':
			runeB = 0x0A
		case 'B', 'b':
			runeB = 0x0B
		case 'C', 'c':
			runeB = 0x0C
		case 'D', 'd':
			runeB = 0x0D
		case 'E', 'e':
			runeB = 0x0E
		case 'F', 'f':
			runeB = 0x0F
		}
		runeA = runeA | runeB
		//buf.WriteByte(temp)
		buf.WriteByte(byte(runeA))
		//fmt.Println(buf.Len())
	}
	//fmt.Println(buf.Len())
	return buf.Bytes()
}

//uuid or node
func HexToString(hex []byte)string{
	//fmt.Println("data:",hex)
	var buf bytes.Buffer
	len := len(hex)
	for i:=0;i<len;i++{
		//处理高四位
		temp := hex[i] >> 4
		if temp <= 0x09{
			var c = 0x30
			t := int(temp)
			c = c + t
			buf.WriteRune(rune(c))
		}else{
			var c = 0x41
			t := int(temp-0x0A)
			c = c + t
			buf.WriteRune(rune(c))
		}
		//处理低四位
		temp = hex[i] & 0x0F
		if temp <= 0x09{
			var c = 0x30
			t := int(temp)
			c = c + t
			buf.WriteRune(rune(c))
		}else{
			var c = 0x41
			t := int(temp-0x0A)
			c = c + t
			buf.WriteRune(rune(c))
		}
	}
	return buf.String()
}

func IPBytesToString(data []byte)(string,error){
	//4 byte
	if len(data) != 4{
		return "",errors.New("the len is less than 4,so it can't be convert")
	}
	var buf bytes.Buffer
	buf.WriteString(strconv.Itoa(int(data[0])))
	buf.WriteRune('.')
	buf.WriteString(strconv.Itoa(int(data[1])))
	buf.WriteRune('.')
	buf.WriteString(strconv.Itoa(int(data[2])))
	buf.WriteRune('.')
	buf.WriteString(strconv.Itoa(int(data[3])))
	return buf.String(),nil
}

func StrPortToBytes(port string)([]byte,error){
	var x uint16
	q,_ := strconv.Atoi(port)
	if q >= 65536{
		return nil,errors.New("the port number is bigger than 65535")
	}
	x = uint16(q)
	w := []byte{uint8(x >> 8),uint8(x)}
	//fmt.Printf("%x",w)
	return w,nil
}

func StrIPToBytes(ip string)([]byte,error){
	strArr := strings.Split(ip,".")
	if len(strArr) != 4{
		return nil, errors.New("the date format is wrong")
	}
	buf := bytes.NewBuffer([]byte{})
	var temp int
	temp,_ = strconv.Atoi(strArr[0])
	a := uint8(temp)
	buf.Write([]byte{a})
	temp,_ = strconv.Atoi(strArr[1])
	a = uint8(temp)
	buf.Write([]byte{a})
	temp,_ = strconv.Atoi(strArr[2])
	a = uint8(temp)
	buf.Write([]byte{a})
	temp,_ = strconv.Atoi(strArr[3])
	a = uint8(temp)
	buf.Write([]byte{a})
	return buf.Bytes(),nil
}

//func Uint8ToByte(v uint8)byte{
//
//}

func PortBytesToInt32(bytes []byte)int32{
	var temp int32
	temp = int32(bytes[0])<<8
	temp = temp + int32(bytes[1])
	return temp
}