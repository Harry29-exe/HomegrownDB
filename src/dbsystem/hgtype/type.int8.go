package hgtype

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/common/random"
	"bytes"
	"strconv"
)

var _ Type = Int8{}

type Int8 struct{}

func (i Int8) Tag() TypeTag {
	return TypeInt8
}

func (i Int8) Skip(args Args, data []byte) []byte {
	return data[8:]
}

func (i Int8) Copy(args Args, dest []byte, data []byte) (copiedBytes int) {
	return copy(dest[:8], data[:8])
}

func (i Int8) IsToastPtr(args Args, data []byte) bool {
	return false
}

func (i Int8) Value(args Args, data []byte) (value []byte) {
	return data[:8]
}

func (i Int8) ValueAndSkip(args Args, data []byte) (value, next []byte) {
	return data[:8], data[8:]
}

func (i Int8) WriteTuple(args Args, dest []byte, value []byte) int {
	return copy(dest, value)
}

func (i Int8) Equal(args Args, v1, v2 []byte) bool {
	return bytes.Equal(v1, v2)
}

func (i Int8) Cmp(args Args, v1, v2 []byte) int {
	return bytes.Compare(v1, v2)
}

func (i Int8) ToStr(args Args, val []byte) string {
	v, _ := bparse.Deserialize.Int8(val)
	return strconv.FormatInt(v, 10)

}

func (i Int8) Rand(args Args, r random.Random) []byte {
	return bparse.Serialize.Int8(r.Int63())
}
