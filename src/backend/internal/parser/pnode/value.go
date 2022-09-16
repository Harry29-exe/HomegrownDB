package pnode

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/column/ctypes"
)

type Value struct {
	V    any
	Type ValueType
}

type ValueType = uint8

const (
	ValueTypeStr ValueType = iota
	ValueTypeInt
	ValueTypeFloat
)

func (v Value) IsAssignableTo(p column.Type) bool {
	switch p {
	case ctypes.Int2:
		return v.Type == ValueTypeInt
	}

	panic("not supported ctype")
}

func (v Value) ConvertTo(ctype column.Type) []byte {
	switch ctype {
	case ctypes.Int2:
		intV := v.V.(int)
		return bparse.Serialize.Int2(int16(intV))
	default:
		panic("not supported ctype")
	}
}
