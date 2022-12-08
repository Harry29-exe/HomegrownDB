package ctype

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/common/random"
	"HomegrownDB/dbsystem/dberr"
	"bytes"
	"strconv"
)

var int8Factory factory = factoryInt8{}

type factoryInt8 struct{}

func (i factoryInt8) Build(args map[string]any) (*CType, dberr.DBError) {
	return newCType(
		TypeInt8,
		int8{},
		int8{},
		int8{},
		int8{},
		false,
		ToastNone,
	), nil
}

var _ Reader = int8{}

type int8 struct{}

func (i int8) Skip(data []byte) []byte {
	return data[8:]
}

func (i int8) Value(data []byte) (value []byte) {
	return data[:8]
}

func (i int8) ValueAndSkip(data []byte) (value, next []byte) {
	return data[:8], data[8:]
}

func (i int8) Copy(dst []byte, data []byte) (copiedBytes int) {
	return copy(dst[:8], data[:8])
}

var _ Writer = int8{}

func (i int8) WriteTuple(dest []byte, value []byte) int {
	return copy(dest, value)
}

func (i int8) WriteNormalized(dest []byte, value []byte) int {
	return copy(dest, value)
}

var _ Operations = int8{}

func (i int8) Equal(v1, v2 []byte) bool {
	return bytes.Equal(v1, v2)
}

func (i int8) Cmp(v1, v2 []byte) int {
	// as long as everywhere binary.BigEndian is used
	// to convert values this is safe
	return bytes.Compare(v1, v2)
}

var _ Debug = int8{}

func (i int8) ToStr(val []byte) string {
	v, _ := bparse.Deserialize.Int8(val)
	return strconv.FormatInt(v, 10)
}

func (i int8) Rand(r random.Random) []byte {
	return bparse.Serialize.Int8(r.Int63())
}
