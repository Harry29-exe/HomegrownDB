package rawtype

import (
	"HomegrownDB/dbsystem/hgtype/typeerr"
	"HomegrownDB/lib/bparse"
	"HomegrownDB/lib/random"
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

func (s Str) Equal(v1, v2 []byte) bool {
	return bytes.Equal(v1, v2)
}

func (s Str) Cmp(v1, v2 []byte) int {
	return bytes.Compare(v1, v2)
}

var _ TypeDebug = Str{}

func (s Str) ToStr(val []byte) string {
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

func (s Str) Validate(args Args, value Value) ValidateResult {
	switch value.TypeTag {
	case TypeStr:
		l := StrUtils.StrLen(value.NormValue)
		if len(value.NormValue) == 0 && !args.Nullable {
			return ValidateResult{Status: ValidateErr, Reason: typeerr.NullNotAllowed{}}
		} else if l > args.Length {
			return ValidateResult{Status: ValidateErr, Reason: typeerr.ToLongErr{}}
		} else if !args.UTF8 && !StrUtils.IsASCII(value.NormValue) {
			return ValidateResult{Status: ValidateErr, Reason: typeerr.UTF8NotAllowed{}}
		}
		return ValidateResult{Status: ValidateOk}
	default:
		return ValidateResult{Status: ValidateErr}
	}
}

func (s Str) WriteValue(writer UniWriter, value Value, args Args) error {
	if args.VarLen {
		return s.varLen.WriteValue(writer, value)
	} else {
		dataLen := s.fourByteLen(value.NormValue) - fourByteLen
		if _, err := writer.Write(value.NormValue[4:]); err != nil {
			return err
		}
		for i := 0; i < int(args.Length)-int(dataLen); i++ {
			if err := writer.WriteByte(' '); err != nil {
				return err
			}
		}
	}
	return nil
}

// -------------------------
//      StrUtils
// -------------------------

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

func (u strUtils) StrLen(val []byte) int {
	if VarLenUtils.IsHeaderOneByte(val[0]) {
		return int(val[0])
	} else {
		return int(bparse.Parse.UInt4(val))
	}
}
