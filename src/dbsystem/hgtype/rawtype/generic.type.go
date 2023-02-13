package rawtype

import (
	"HomegrownDB/dbsystem/hgtype/typeerr"
	"HomegrownDB/lib/random"
)

var TypeOp = typeOp{}

type typeOp struct{}

func (t typeOp) Tag() Tag {
	//TODO implement me
	panic("implement me")
}

func (t typeOp) Validate(args Args, value Value) ValidateResult {
	if !args.Nullable && value.NormValue == nil {
		return ValidateResult{Status: ValidateErr, Reason: typeerr.NullNotAllowed{}}
	}

	if (!args.VarLen && args.Length != len(value.NormValue)) ||
		args.VarLen && args.Length < VarLenHelper.dataLen(value.NormValue) {
		return ValidateResult{Status: ValidateErr, Reason: typeerr.ToLongErr{}}
	}

	if !args.UTF8 && value.TypeTag == TypeStr {
		var data []byte
		if args.VarLen {
			data = VarLenHelper.data(value.NormValue)
		} else {
			data = value.NormValue
		}
		for i := 0; i < len(data); i++ {
			if data[i] > 127 {
				return ValidateResult{Status: ValidateErr, Reason: typeerr.UTF8NotAllowed{}}
			}
		}
	}

	return ValidateResult{Status: ValidateOk}
}

// DataLen returns number of bytes without length header
// value is slice that at index 0 has first byte of value, it can contain additional bytes after value
func (t typeOp) DataLen(value []byte, args Args) int {
	if args.VarLen {
		return VarLenHelper.dataLen(value)
	}
	return args.Length
}

// FullLen returns number of bytes that value occupies
func (t typeOp) FullLen(value []byte, args Args) int {
	if args.VarLen {
		return VarLenHelper.fullLen(value)
	}
	return args.Length
}

// GetData returns only data without header
func (t typeOp) GetData(value []byte, args Args) []byte {
	if args.VarLen {
		if VarLenHelper.lenIsOneByte(value[0]) {
			return value[1:VarLenHelper.getOneByteLen(value)]
		}
		return value[4:VarLenHelper.getFourByteLen(value)]
	}
	return value[:args.Length]
}

func (t typeOp) Skip(data []byte, args Args) []byte {
	if args.VarLen {
		return data[VarLenHelper.fullLen(data):]
	}
	return data[args.Length:]
}

func (t typeOp) Copy(dest []byte, data []byte) (copiedBytes int) {
	//TODO implement me
	panic("implement me")
}

func (t typeOp) IsToastPtr(data []byte) bool {
	//TODO implement me
	panic("implement me")
}

func (t typeOp) Value(data []byte) (value []byte) {
	//TODO implement me
	panic("implement me")
}

func (t typeOp) ValueAndSkip(data []byte) (value, next []byte) {
	//TODO implement me
	panic("implement me")
}

func (t typeOp) WriteValueToTuple(tupleWriter UniWriter, value []byte, args Args) error {
	//TODO implement me
	panic("implement me")
}

func (t typeOp) Equal(v1, v2 []byte) bool {
	//TODO implement me
	panic("implement me")
}

func (t typeOp) Cmp(v1, v2 []byte) int {
	//TODO implement me
	panic("implement me")
}

func (t typeOp) ToStr(val []byte) string {
	//TODO implement me
	panic("implement me")
}

func (t typeOp) Rand(args Args, r random.Random) []byte {
	//TODO implement me
	panic("implement me")
}
