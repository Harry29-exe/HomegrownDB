package bparse

import "encoding/binary"

type Serializer struct {
	buffer []byte
}

func NewSerializer() *Serializer {
	return &Serializer{
		[]byte{},
	}
}

func (s *Serializer) GetBytes() []byte {
	return s.buffer
}

func (s *Serializer) Append(bytes []byte) {
	s.buffer = append(s.buffer, bytes...)
}

func (s *Serializer) SmString(str string) {
	strLen := len(str)
	if strLen > 255 {
		panic("SmString can append string with length up to 255")
	}

	s.Uint8(uint8(strLen))
	s.buffer = append(s.buffer, []byte(str)...)
}

func (s *Serializer) MdString(str string) {
	strLen := len(str)
	if strLen > 65536 {
		panic("MdString can append string with length up to 65536")
	}

	s.Uint16(uint16(strLen))
	s.buffer = append(s.buffer, []byte(str)...)
}

func (s *Serializer) Bool(boolean bool) {
	var data uint8 = 0
	if boolean {
		data = 1
	}
	s.buffer = append(s.buffer, data)
}

func (s *Serializer) Uint8(integer uint8) {
	s.buffer = append(s.buffer, integer)
}

func (s *Serializer) Uint16(integer uint16) {
	data := make([]byte, 2)
	binary.BigEndian.PutUint16(data, integer)

	s.buffer = append(s.buffer, data...)
}

func (s *Serializer) Uint32(integer uint32) {
	data := make([]byte, 4)
	binary.BigEndian.PutUint32(data, integer)

	s.buffer = append(s.buffer, data...)
}

func (s *Serializer) Uint64(integer uint64) {
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, integer)

	s.buffer = append(s.buffer, data...)
}

func (s *Serializer) Int64(integer int64) {
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, uint64(integer))

	s.buffer = append(s.buffer, data...)
}
