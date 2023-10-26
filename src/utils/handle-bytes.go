package utils

import (
	"bytes"
	"encoding/binary"
	"reflect"
)

func StringToBytes(s string, size int) []byte {
	var buf bytes.Buffer

	for i := 0; i < size; i++ {
		if i < len(s) {
			buf.WriteByte(s[i])
		} else {
			buf.WriteByte(0x20)
		}
	}

	return buf.Bytes()
}

func EncodeToBytes(p interface{}) []byte {

	buf := bytes.Buffer{}
	values := reflect.ValueOf(p)
	for i := 0; i < values.NumField(); i++ {
		err := binary.Write(&buf, binary.LittleEndian, values.Field(i).Interface())
		if err != nil {
			panic("Error encoding ID:" + err.Error())
		}
	}

	return buf.Bytes()
}
