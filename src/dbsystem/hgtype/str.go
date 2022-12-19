package hgtype

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/common/random"
	"HomegrownDB/dbsystem/hgtype/basetype"
	"HomegrownDB/dbsystem/hgtype/toast"
	"bytes"
)

type StrArgs struct {
	Length     uint32
	VaryingLen bool
	UTF8       bool
}

func NewStr(args StrArgs) HGType {
	ctype := str{
		VarLenType: basetype.VarLenType{
			Length:      args.Length,
			VaryingLen:  args.VaryingLen,
			ToastStatus: toast.PLAIN,
		},
		UTF8: args.UTF8,
	}

	return newHGType(TypeStr, &ctype, &ctype, &ctype, &ctype, true, toast.EXTENDED)
}

type str struct {
	basetype.VarLenType
	UTF8 bool
}

var (
	_ Reader = &str{}
	_ Writer = &str{}
)

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
