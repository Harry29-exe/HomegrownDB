package rawtype

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/common/random"
	"HomegrownDB/dbsystem/hgtype/typeerr"
	"bytes"
	"strconv"
)

var _ Type = Int8{}

type Int8 struct{}

func (i Int8) Tag() Tag {
	return TypeInt8
}

func (i Int8) Skip(data []byte) []byte {
	return data[8:]
}

func (i Int8) Copy(dest []byte, data []byte) (copiedBytes int) {
	return copy(dest[:8], data[:8])
}

func (i Int8) IsToastPtr(data []byte) bool {
	return false
}

func (i Int8) Value(data []byte) (value []byte) {
	return data[:8]
}

func (i Int8) ValueAndSkip(data []byte) (value, next []byte) {
	return data[:8], data[8:]
}

func (i Int8) Validate(args Args, value Value) ValidateResult {
	switch value.TypeTag {
	case TypeInt8:
		return ValidateResult{Status: ValidateOk}
	default:
		return ValidateResult{Status: ValidateErr, Reason: typeerr.TypesNotConvertable{}}
	}
}

func (i Int8) WriteValue(writer UniWriter, value Value, _ Args) error {
	_, err := writer.Write(value.NormValue)
	return err
}

func (i Int8) Equal(v1, v2 []byte) bool {
	return bytes.Equal(v1, v2)
}

func (i Int8) Cmp(v1, v2 []byte) int {
	return bytes.Compare(v1, v2)
}

func (i Int8) ToStr(val []byte) string {
	v, _ := bparse.Deserialize.Int8(val)
	return strconv.FormatInt(v, 10)

}

func (i Int8) Rand(args Args, r random.Random) []byte {
	return bparse.Serialize.Int8(r.Int63())
}
