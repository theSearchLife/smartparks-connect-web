package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"math"
	"strings"
)

type VType string

var VTypes = []VType{"uint32", "uint8", "bool", "byte_array", "string", "float"}

func ConvertBytes(idHex string, vtype VType, length uint8, content interface{}) (data []byte, err error) {
	idByte, err := hexSingleDecode(idHex)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer([]byte{idByte})
	buf.WriteByte(length)
	if length == 0 {
		return buf.Bytes(), nil
	}
	switch vtype {
	case "int32":
		buf.Write(Int32ToBytes(int32(content.(float64))))
	case "uint32":
		buf.Write(Uint32ToBytes(uint32(content.(float64))))
	case "uint8":
		buf.WriteByte(uint8(content.(float64)))
	case "bool":
		if content.(string) == "true" {
			buf.WriteByte(1)
		} else {
			buf.WriteByte(0)
		}
	case "byte_array":
		bs, err := hexDecode(content.(string))
		if err != nil {
			return nil, err
		}
		buf.Write(bs)
	case "string":
		buf.Write([]byte(content.(string)))
	case "float":
		buf.Write(Uint32ToBytes(math.Float32bits(float32(content.(float64)))))
	default:
		return nil, errors.New("unknow variable type")
	}
	return buf.Bytes(), nil
}
func hexSingleDecode(hexString string) (byte, error) {
	if strings.HasPrefix(hexString, "0x") {
		hexString = hexString[2:]
	}
	bs, err := hex.DecodeString(hexString)
	if err != nil {
		return 0, err
	}
	return bs[0], nil
}
func hexDecode(hexString string) ([]byte, error) {
	return hex.DecodeString(hexString)
}

func Uint32ToBytes(num uint32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, num)
	return buf
}
func Int32ToBytes(num int32) []byte {
	buff := new(bytes.Buffer)
	_ = binary.Write(buff, binary.LittleEndian, num)

	return buff.Bytes()
}
