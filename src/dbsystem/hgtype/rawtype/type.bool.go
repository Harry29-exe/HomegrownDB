package rawtype

import (
	"HomegrownDB/dbsystem/hgtype/typeerr"
	"HomegrownDB/lib/random"
)

type Bool struct{}

func (b Bool) Tag() Tag {
	return TypeBool
}

func (b Bool) Validate(args Args, value Value) ValidateResult {
	if !args.Nullable && value.NormValue == nil {
		return ValidateResult{Status: ValidateErr, Reason: typeerr.TypesNotConvertable{}}
	} else {
		return ValidateResult{Status: ValidateOk}
	}
}

func (b Bool) Skip(data []byte) []byte {
	return data[1:]
}

func (b Bool) Copy(dest []byte, data []byte) (copiedBytes int) {
	dest[0] = data[0]
	return 1
}

func (b Bool) IsToastPtr(data []byte) bool {
	return false
}

func (b Bool) Value(data []byte) (value []byte) {
	return data[0:1]
}

func (b Bool) ValueAndSkip(data []byte) (value, next []byte) {
	return data[0:1], data[1:]
}

func (b Bool) WriteValue(writer UniWriter, value Value, args Args) error {
	return writer.WriteByte(value.NormValue[0])
}

func (b Bool) Equal(v1, v2 []byte) bool {
	return v1[0] == v2[0]
}

func (b Bool) Cmp(v1, v2 []byte) int {
	switch {
	case v1[0] == v2[0]:
		return 0
	case v1[0] < v2[0]:
		return -1
	default:
		return 1
	}
}

func (b Bool) ToStr(val []byte) string {
	if val[0] == 0 {
		return "false"
	}
	return "true"
}

func (b Bool) Rand(args Args, r random.Random) []byte {
	//TODO implement me
	panic("implement me")
}

var _ Type = Bool{}
