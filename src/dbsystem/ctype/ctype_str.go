package ctype

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/common/random"
	"HomegrownDB/dbsystem/ctype/toast"
	"HomegrownDB/dbsystem/dberr"
	"bytes"
)

var strFactory factory = factoryStr{}

type factoryStr struct{}

func (f factoryStr) Build(args map[string]any) (*CType, dberr.DBError) {
	ctype := str{}

	length, ok := args["length"]
	if !ok {
		length = 1
	} else {
		switch v := length.(type) {
		case uint32:
			ctype.Length = v
		default:
			return nil, NewFactoryArgError(CTypeToStr(TypeInt8), "length", "uint32")
		}
	}

	varyingLen, ok := args["varyingLength"]
	if !ok {
		varyingLen = false
	} else {
		switch v := varyingLen.(type) {
		case bool:
			ctype.VaryingLength = v
		default:
			return nil, NewFactoryArgError(CTypeToStr(TypeInt8), "varyingLength", "bool")
		}
	}

	utf8, ok := args["utf8"]
	if !ok {
		utf8 = true
	} else {
		switch v := utf8.(type) {
		case bool:
			ctype.UTF8 = v
		default:
			return nil, NewFactoryArgError(CTypeToStr(TypeInt8), "utf8", "bool")
		}
	}

	return newCType(TypeStr, &ctype, &ctype, &ctype, &ctype, true, ToastStore), nil
}

type str struct {
	Length        uint32
	VaryingLength bool
	UTF8          bool
}

var _ Reader = &str{}

func (s *str) Skip(data []byte) []byte {
	if toast.IsToasted(data[0]) {
		return data[toast.InTupleSize:]
	}

	if s.lenIsOneByte(data[0]) {
		l := s.oneByteLen(data)
		return data[l:]
	} else {
		l := s.fourByteLen(data)
		return data[l:]
	}
}
func (s *str) Copy(dest []byte, data []byte) int {
	if toast.IsToasted(data[0]) {
		return copy(dest, data[:toast.InTupleSize])
	}

	if s.lenIsOneByte(data[0]) {
		l := s.oneByteLen(data)
		return copy(dest, data[:l])
	} else {
		l := s.fourByteLen(data)
		return copy(dest, data[:l+1])
	}
}

func (s *str) Value(data []byte) []byte {
	if s.lenIsOneByte(data[0]) {
		l := s.oneByteLen(data)
		val := make([]byte, l+3)
		bparse.Serialize.PutUInt4(uint32(l)+3, val)
		copy(val[4:], data[1:l])

		return val
	} else {
		l := s.fourByteLen(data)
		val := make([]byte, l)
		copy(val, data[:l])
		return val
	}
}

func (s *str) ValueAndSkip(data []byte) (value, next []byte) {
	if s.lenIsOneByte(data[0]) { // 1 byte header
		l := s.oneByteLen(data)
		next = data[l:]
		val := make([]byte, l+3)
		bparse.Serialize.PutUInt4(uint32(l)+3, val)
		copy(val[4:], data[1:l])

		return val, next
	} else {
		l := s.fourByteLen(data)
		value = make([]byte, l)
		copy(value, data[:l])
		next = data[l:]
		return value, next
	}
}

var _ Writer = &str{}

func (s *str) WriteTuple(dest []byte, value []byte) int {
	if toast.IsToasted(value[0]) {
		return copy(dest, value[:toast.InTupleSize])
	}

	if s.lenIsOneByte(value[0]) { // 1 byte header
		l := s.oneByteLen(value)
		return copy(dest, value[:l])
	}

	l := s.fourByteLen(value)
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

func (s *str) WriteNormalized(dest []byte, value []byte) int {
	//TODO implement me
	panic("implement me")
}

var _ Operations = &str{}

func (s *str) Equal(v1, v2 []byte) bool {
	return bytes.Equal(v1, v2)
}

func (s *str) Cmp(v1, v2 []byte) int {
	return bytes.Compare(v1, v2)
}

var _ Debug = &str{}

func (s *str) ToStr(val []byte) string {
	return string(val)
}

func (s *str) Rand(r random.Random) []byte {
	buff := bytes.Buffer{}
	l := r.Int64mm(0, int64(s.Length))
	if l > 127 {
		l = 127
	}
	buff.WriteByte(byte(l))
	for i := 0; i < int(l); i++ {
		buff.WriteByte(byte(r.CharASCII()))
	}

	return buff.Bytes()
}

func (s *str) lenIsOneByte(firstByte byte) bool {
	return firstByte > 127
}

func (s *str) fourByteLen(data []byte) uint32 {
	return bparse.Parse.UInt4(data) & strFourByteHeaderMask
}

func (s *str) oneByteLen(data []byte) uint8 {
	return data[0] & strOneByteHeaderMask
}

// 01111111
var strOneByteHeaderMask = byte(127)

// 00111111 11111111 11111111 11111111
var strFourByteHeaderMask = uint32(1073741823)
