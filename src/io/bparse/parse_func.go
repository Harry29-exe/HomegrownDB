package bparse

import "encoding/binary"

type parser struct{}

var Parse = parser{}

func (p parser) SmString(data []byte) (value string) {
	return string(data[1 : uint16(data[0])+1])
}

func (p parser) MdString(data []byte) (value string) {
	return string(data[2 : binary.LittleEndian.Uint16(data)+2])
}

func (p parser) Int2(data []byte) (value int16) {
	return int16(binary.LittleEndian.Uint16(data[:2]))
}

func (p parser) UInt2(data []byte) (value uint16) {
	return binary.LittleEndian.Uint16(data[:2])
}

func (p parser) Int4(data []byte) (value int32) {
	return int32(binary.LittleEndian.Uint32(data[:4]))
}

func (p parser) UInt4(data []byte) (value uint32) {
	return binary.LittleEndian.Uint32(data[:4])
}

func (p parser) Int8(data []byte) (value int64) {
	return int64(binary.LittleEndian.Uint64(data[:8]))
}

func (p parser) SmBytes(data []byte) (value []byte) {
	return data[1 : uint16(data[0])+2]
}
