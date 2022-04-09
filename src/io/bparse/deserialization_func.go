package bparse

import "encoding/binary"

type deserializer struct{}

var Deserialize = deserializer{}

func (d deserializer) SmString(data []byte) (value string, subsequent []byte) {
	strLen := int(data[0])

	endPointer := 1 + strLen
	return string(data[1:endPointer]), data[endPointer:]
}

func (d deserializer) MdString(data []byte) (value string, subsequent []byte) {
	strLen := binary.LittleEndian.Uint16(data)
	afterString := strLen + 2
	return string(data[2:afterString]), data[afterString:]
}

func (d deserializer) Int2(data []byte) (value int16, subsequent []byte) {
	value = int16(binary.LittleEndian.Uint16(data[:2]))
	subsequent = data[2:]
	return
}
