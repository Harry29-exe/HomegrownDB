package bparse

import "encoding/binary"

type serializer struct{}

var Serialize serializer = serializer{}

func (s *serializer) Int2(data []byte) (value int16, subsequent []byte) {
	value = int16(binary.LittleEndian.Uint16(data[:2]))
	subsequent = data[2:]
	return
}
