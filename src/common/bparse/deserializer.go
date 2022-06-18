package bparse

import "encoding/binary"

type Deserializer struct {
	data    []byte
	pointer uint64
}

func NewDeserializer(data []byte) *Deserializer {
	return &Deserializer{
		data:    data,
		pointer: 0,
	}
}

// SmString read string with 1 byte length prefix
func (d *Deserializer) SmString() string {
	strLen := d.data[d.pointer]
	d.pointer++

	pointer := d.pointer
	d.pointer = d.pointer + uint64(strLen)
	return string(d.data[pointer:d.pointer])
}

// MdString read string with 2 bytes length prefix
func (d *Deserializer) MdString() string {
	strLen := binary.LittleEndian.Uint16(d.data[d.pointer : d.pointer+2])
	d.pointer += 2

	pointer := d.pointer
	d.pointer = d.pointer + uint64(strLen)
	return string(d.data[pointer:d.pointer])
}

func (d *Deserializer) Bool() bool {
	pointer := d.pointer
	d.pointer++
	return uint8(d.data[pointer]) > 0
}

func (d *Deserializer) Uint8() uint8 {
	pointer := d.pointer
	d.pointer++

	return d.data[pointer]
}

func (d *Deserializer) Uint16() uint16 {
	pointer := d.pointer
	d.pointer += 2

	return binary.LittleEndian.Uint16(d.data[pointer:d.pointer])
}

func (d *Deserializer) Uint32() uint32 {
	pointer := d.pointer
	d.pointer += 4

	return binary.LittleEndian.Uint32(d.data[pointer:d.pointer])
}

func (d *Deserializer) Uint64() uint64 {
	pointer := d.pointer
	d.pointer += 8

	return binary.LittleEndian.Uint64(d.data[pointer:d.pointer])
}

func (d *Deserializer) Int16() int16 {
	pointer := d.pointer
	d.pointer += 2

	return int16(binary.LittleEndian.Uint16(d.data[pointer:d.pointer]))
}

func (d *Deserializer) Int32() int32 {
	pointer := d.pointer
	d.pointer += 4

	return int32(binary.LittleEndian.Uint32(d.data[pointer:d.pointer]))
}

func (d *Deserializer) Int64() int64 {
	pointer := d.pointer
	d.pointer += 8

	return int64(binary.LittleEndian.Uint64(d.data[pointer:d.pointer]))
}

func (d *Deserializer) IsEmpty() bool {
	return d.pointer >= uint64(len(d.data))
}

func (d *Deserializer) RemainedData() []byte {
	return d.data[d.pointer:]
}
