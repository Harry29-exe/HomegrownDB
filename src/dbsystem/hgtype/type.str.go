package hgtype

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/common/random"
	"bytes"
	"unicode"
)

var _ Type = Str{}

type Str struct {
	varLen
}

func (s Str) Tag() Tag {
	return TypeStr
}

var _ TypeOperations = Str{}

func (s Str) Equal(args Args, v1, v2 []byte) bool {
	return bytes.Equal(v1, v2)
}

func (s Str) Cmp(args Args, v1, v2 []byte) int {
	return bytes.Compare(v1, v2)
}

var _ TypeDebug = Str{}

func (s Str) ToStr(args Args, val []byte) string {
	return string(val)
}

func (s Str) Rand(args Args, r random.Random) []byte {
	buff := bytes.Buffer{}
	l := r.Int64mm(0, int64(args.Length))
	if l > 127 {
		l = 127
	}
	buff.WriteByte(byte(l))
	for i := 0; i < int(l); i++ {
		buff.WriteByte(byte(r.CharASCII()))
	}

	return buff.Bytes()
}

var StrUtils = strUtils{}

type strUtils struct{}

func (u strUtils) IsASCII(val []byte) bool {
	var data []byte
	if VarLenUtils.IsHeaderOneByte(val[0]) {
		data = val[1:val[0]]
	} else {
		l := bparse.Parse.UInt4(val)
		data = val[4:l]
	}

	for i := 0; i < len(data); i++ {
		if data[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func (u strUtils) StrLen(val []byte) uint32 {
	if VarLenUtils.IsHeaderOneByte(val[0]) {
		return uint32(val[0])
	} else {
		return bparse.Parse.UInt4(val)
	}
}
