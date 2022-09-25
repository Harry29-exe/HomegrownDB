package pnode

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/column/ctypes"
	"fmt"
	"strconv"
)

type Value interface {
	Node() Node
	V() any
	VasStr() string
	Type() ValueType
	TypeAsStr() string

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
	N Node
	v int
}

var intValueSupportedCTypes = map[column.Type]bool{
	ctypes.Int2: true,
}

func (i IntValue) Node() Node {
	return i.N
}

func (i IntValue) V() any {
	return i.v
}

func (i IntValue) VasStr() string {
	return strconv.Itoa(i.v)
}

func (i IntValue) Type() ValueType {
	return ValueTypeInt
}

func (i IntValue) TypeAsStr() string {
	return "Int"
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
	N Node
	v float64
}

var floatValueSupportedCTypes = map[column.Type]bool{}

func (i FloatValue) Node() Node {
	return i.N
}

func (i FloatValue) V() any {
	return i.v
}

func (i FloatValue) VasStr() string {
	return fmt.Sprintf("%f", i.v)
}

func (i FloatValue) Type() ValueType {
	return ValueTypeFloat
}

func (i FloatValue) TypeAsStr() string {
	return "Float"
}

func (i FloatValue) IsAssignableTo(ctype column.Type) bool {
	return floatValueSupportedCTypes[ctype]
}

func (i FloatValue) ConvertTo(ctype column.Type) []byte {
	switch ctype {
	default:
		panic("not supported type")
	}
}

// ### StrValue

type StrValue struct {
	N Node
	v string
}

func (i StrValue) Node() Node {
	return i.N
}

var strValueSupportedCTypes = map[column.Type]bool{}

func (i StrValue) V() any {
	return i.v
}

func (i StrValue) VasStr() string {
	return i.v
}

func (i StrValue) Type() ValueType {
	return ValueTypeStr
}

func (i StrValue) TypeAsStr() string {
	return "String"
}

func (i StrValue) IsAssignableTo(ctype column.Type) bool {
	return strValueSupportedCTypes[ctype]
}

func (i StrValue) ConvertTo(ctype column.Type) []byte {
	switch ctype {
	default:
		panic("not supported type")
	}
}
