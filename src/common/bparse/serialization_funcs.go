package bparse

import (
	"encoding/binary"
	"math"
)

var Serialize = serializer{}

type serializer struct{}

func (s serializer) Int2(value int16) []byte {
	bytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes, uint16(value))

	return bytes
}

func (s serializer) PutInt2(value int16, dest []byte) {
	binary.LittleEndian.PutUint16(dest, uint16(value))
}

func (s serializer) Uint2(value uint16) []byte {
	bytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes, value)

	return bytes
}

func (s serializer) PutUInt2(value uint16, dest []byte) {
	binary.LittleEndian.PutUint16(dest, value)
}

func (s serializer) Int4(value int32) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, uint32(value))

	return bytes
}

func (s serializer) Uint4(value uint32) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, value)

	return bytes
}

func (s serializer) Int8(value int64) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, uint64(value))

	return bytes
}

func (s serializer) Uint8(value uint64) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, value)

	return bytes
}

func (s serializer) Float8(value float64) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, math.Float64bits(value))

	return bytes
}
