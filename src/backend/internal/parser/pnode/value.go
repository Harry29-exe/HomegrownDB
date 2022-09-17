package pnode

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/column/ctypes"
)

type Value interface {
	V() any
	Type() ValueType
	IsAssignableTo(ctype column.Type) bool
	ConvertTo(ctype column.Type) []byte
}

type ValueType = uint8

const (
	ValueTypeStr ValueType = iota
	ValueTypeInt
	ValueTypeFloat
)

// ### IntValue ###

type IntValue struct {
	v int
}

var intValueSupportedCTypes = map[column.Type]bool{
	ctypes.Int2: true,
}

func (i IntValue) V() any {
	return i.v
}

func (i IntValue) Type() ValueType {
	return ValueTypeInt
}

func (i IntValue) IsAssignableTo(ctype column.Type) bool {
	return intValueSupportedCTypes[ctype]
}

func (i IntValue) ConvertTo(ctype column.Type) []byte {
	switch ctype {
	case ctypes.Int2:
		return bparse.Serialize.Int2(int16(i.v))
	default:
		panic("not supported type")
	}
}

// ### FloatValue

type FloatValue struct {
	v float64
}

var floatValueSupportedCTypes = map[column.Type]bool{}

func (i FloatValue) V() any {
	return i.v
}

func (i FloatValue) Type() ValueType {
	return ValueTypeFloat
}

func (i FloatValue) IsAssignableTo(ctype column.Type) bool {
	return intValueSupportedCTypes[ctype]
}

func (i FloatValue) ConvertTo(ctype column.Type) []byte {
	switch ctype {
	default:
		panic("not supported type")
	}
}

// ### StrValue

type StrValue struct {
	v string
}

var strValueSupportedCTypes = map[column.Type]bool{}

func (i StrValue) V() any {
	return i.v
}

func (i StrValue) Type() ValueType {
	return ValueTypeFloat
}

func (i StrValue) IsAssignableTo(ctype column.Type) bool {
	return intValueSupportedCTypes[ctype]
}

func (i StrValue) ConvertTo(ctype column.Type) []byte {
	switch ctype {
	default:
		panic("not supported type")
	}
}
