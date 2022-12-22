package hgtype

import (
	"HomegrownDB/common/random"
	"bytes"
)

var _ Type = Str{}

type Str struct {
	varLen
}

func (s Str) Tag() TypeTag {
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
