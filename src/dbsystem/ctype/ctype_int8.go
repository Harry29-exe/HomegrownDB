package ctype

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/common/random"
	"bytes"
	"strconv"
)

var _ CType = &Int8{}

type Int8 struct {
}

func (i Int8) Skip(data []byte) []byte {
	return data[8:]
}

func (i Int8) Value(data []byte) (value []byte) {
	return data[:8]
}

func (i Int8) ValueAndSkip(data []byte) (value, next []byte) {
	return data[:8], data[8:]
}

func (i Int8) Copy(dst []byte, data []byte) (copiedBytes int) {
	return copy(dst[:8], data[:8])
}

func (i Int8) Equal(v1, v2 []byte) bool {
	return bytes.Equal(v1, v2)
}

func (i Int8) Cmp(v1, v2 []byte) int {
	// as long as everywhere binary.BigEndian is used
	// to convert values this is safe
	return bytes.Compare(v1, v2)
}

func (i Int8) ToStr(val []byte) string {
	v, _ := bparse.Deserialize.Int8(val)
	return strconv.FormatInt(v, 10)
}

func (i Int8) Rand(r random.Random) []byte {
	return bparse.Serialize.Int8(r.Int63())
}
