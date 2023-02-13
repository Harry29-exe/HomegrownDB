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
	strLen := binary.BigEndian.Uint16(data)
	afterString := strLen + 2
	return string(data[2:afterString]), data[afterString:]
}

func (d deserializer) UInt2(data []byte) (value uint16, subsequent []byte) {
	value = binary.BigEndian.Uint16(data[:2])
	subsequent = data[2:]
	return
}

func (d deserializer) Int2(data []byte) (value int16, subsequent []byte) {
	value = int16(binary.BigEndian.Uint16(data[:2]))
	subsequent = data[2:]
	return
}

func (d deserializer) Int4(data []byte) (value int32, subsequent []byte) {
	value = int32(binary.BigEndian.Uint32(data[:4]))
	subsequent = data[4:]
	return
}

func (d deserializer) UInt4(data []byte) (value uint32, subsequent []byte) {
	value = binary.BigEndian.Uint32(data[:4])
	subsequent = data[4:]
	return
}

func (d deserializer) Int8(data []byte) (value int64, subsequent []byte) {
	value = int64(binary.BigEndian.Uint64(data[:8]))
	subsequent = data[8:]
	return
}

func (d deserializer) SmBytes(data []byte) (value []byte, subsequent []byte) {
	indexAfterVal := uint16(data[0]) + 2
	value = data[1:indexAfterVal]
	subsequent = data[indexAfterVal:]
	return
}
