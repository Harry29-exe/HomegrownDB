package hgtype

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/hgtype/toast"
	"math"
)

const (
	UnknownVarLenSize = int(math.MaxUint32 & toast.FourByteHeaderMask)
	MaxVarLenSize     = int(math.MaxUint32&toast.FourByteHeaderMask) - 1
)

var (
	_ TypeReader = varLen{}
	_ TypeWriter = varLen{}
)

type varLen struct{}

// -------------------------
//      TypeReader
// -------------------------

func (v varLen) Skip(data []byte) []byte {
	if toast.IsVarLenToasted(data[0]) {
		return data[:toast.InTupleSize]
	}

	if v.lenIsOneByte(data[0]) {
		l := v.oneByteLen(data)
		return data[l:]
	} else {
		l := v.fourByteLen(data)
		return data[l:]
	}
}

func (v varLen) Copy(dest []byte, data []byte) (copiedBytes int) {
	if toast.IsVarLenToasted(data[0]) {
		return copy(dest, data[:toast.InTupleSize])
	}

	if v.lenIsOneByte(data[0]) {
		l := v.oneByteLen(data)
		return copy(dest, data[:l])
	} else {
		l := v.fourByteLen(data)
		return copy(dest, data[:l+1])
	}
}

func (v varLen) IsToastPtr(data []byte) bool {
	return toast.IsVarLenToasted(data[0])
}

func (v varLen) Value(data []byte) (value []byte) {
	if v.lenIsOneByte(data[0]) {
		l := v.oneByteLen(data)
		val := make([]byte, l+3)
		bparse.Serialize.PutUInt4(uint32(l)+3, val)
		copy(val[4:], data[1:l])

		return val
	} else {
		l := v.fourByteLen(data)
		val := make([]byte, l)
		copy(val, data[:l])
		return val
	}
}

func (v varLen) ValueAndSkip(data []byte) (value, next []byte) {
	if v.lenIsOneByte(data[0]) { // 1 byte header
		l := v.oneByteLen(data)
		next = data[l:]
		val := make([]byte, l+3)
		bparse.Serialize.PutUInt4(uint32(l)+3, val)
		copy(val[4:], data[1:l])

		return val, next
	} else {
		l := v.fourByteLen(data)
		value = make([]byte, l)
		copy(value, data[:l])
		next = data[l:]
		return value, next
	}
}

// -------------------------
//      TypeWriter
// -------------------------

func (v varLen) WriteTuple(dest []byte, value []byte) int {
	//TODO implement me
	panic("implement me")
}

// -------------------------
//      private
// -------------------------

func (v varLen) lenIsOneByte(firstByte byte) bool {
	return firstByte > 127
}

func (v varLen) fourByteLen(data []byte) uint32 {
	return bparse.Parse.UInt4(data) & fourByteHeaderMask
}

func (v varLen) oneByteLen(data []byte) uint8 {
	return data[0] & oneByteHeaderMask
}

// 01111111
var oneByteHeaderMask = byte(127)

// 00111111 11111111 11111111 11111111
var fourByteHeaderMask = uint32(1073741823)

var VarLenUtils = varLenUtils{}

type varLenUtils struct{}

func (v varLenUtils) IsHeaderOneByte(firstByte byte) bool {
	return firstByte > 127
}
