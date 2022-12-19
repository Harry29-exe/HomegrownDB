package basetype

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/hgtype/toast"
)

type VarLenType struct {
	Length      uint32
	VaryingLen  bool
	ToastStatus toast.Status
}

// -------------------------
//      hgtype.Reader implementation
// -------------------------

func (v VarLenType) Skip(data []byte) []byte {
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

func (v VarLenType) Copy(dest []byte, data []byte) (copiedBytes int) {
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

func (v VarLenType) IsToastPtr(data []byte) bool {
	return toast.IsVarLenToasted(data[0])
}

func (v VarLenType) Value(data []byte) (value []byte) {
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

func (v VarLenType) ValueAndSkip(data []byte) (value, next []byte) {
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
//      hgtype.Writer implementation
// -------------------------

func (v VarLenType) WriteTuple(dest []byte, value []byte) int {
	if toast.IsVarLenToasted(value[0]) {
		return copy(dest, value[:toast.InTupleSize])
	}

	if v.lenIsOneByte(value[0]) { // 1 byte header
		l := v.oneByteLen(value)
		return copy(dest, value[:l])
	}

	l := v.fourByteLen(value)
	if toast.IsCompressed(value[0]) {
		//todo implement me
		panic("Not implemented")
	}
	if l <= 126+4 {
		dest[0] = uint8(l - 3 + 128)
		return copy(dest[1:], value[4:l]) + 1
	}
	return copy(dest, value[:l])
}

func (v VarLenType) WriteNormalized(dest []byte, value []byte) int {
	//TODO implement me
	panic("implement me")
}

func (v VarLenType) lenIsOneByte(firstByte byte) bool {
	return firstByte > 127
}

func (v VarLenType) fourByteLen(data []byte) uint32 {
	return bparse.Parse.UInt4(data) & fourByteHeaderMask
}

func (v VarLenType) oneByteLen(data []byte) uint8 {
	return data[0] & oneByteHeaderMask
}

// 01111111
var oneByteHeaderMask = byte(127)

// 00111111 11111111 11111111 11111111
var fourByteHeaderMask = uint32(1073741823)
