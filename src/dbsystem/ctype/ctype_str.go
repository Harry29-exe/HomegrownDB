package ctype

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/common/random"
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

	return newCType(&ctype, &ctype, &ctype, true, ToastStore), nil
}

var (
	_ Reader     = &str{}
	_ Operations = &str{}
	_ Debug      = &str{}
)

type str struct {
	Length        uint32
	VaryingLength bool
	UTF8          bool
}

func (s *str) Skip(data []byte) []byte {
	if data[0] > 127 { // toast - in bg table
		l := bparse.Parse.UInt2(data)
		return data[l+2:]
	} else {
		l := data[0]
		return data[l+1:]
	}
}

func (s *str) Value(data []byte) []byte {
	if data[0] > 127 { // toast - in bg table
		l := bparse.Parse.UInt2(data)
		return data[2 : l+2]
	} else {
		l := data[0]
		return data[1 : l+1]
	}
}

func (s *str) ValueAndSkip(data []byte) (value, next []byte) {
	if data[0] > 127 { // toast - in bg table
		l := bparse.Parse.UInt2(data)
		value = data[2 : l+2]
		next = data[l+2:]
	} else {
		l := data[0]
		value = data[1 : l+1]
		next = data[l+1:]
	}
	return
}

func (s *str) Copy(dest []byte, data []byte) int {
	if data[0] > 127 {
		l := bparse.Parse.UInt2(data)
		return copy(dest, data[:l+2])
	} else {
		l := data[0]
		return copy(dest, data[:l+1])
	}
}

func (s *str) Equal(v1, v2 []byte) bool {
	return bytes.Equal(v1, v2)
}

func (s *str) Cmp(v1, v2 []byte) int {
	return bytes.Compare(v1, v2)
}

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
