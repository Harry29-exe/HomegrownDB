package intype

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/hgtype/rawtype"
)

func ConvInt8(val int64) []byte {
	return bparse.Serialize.Int8(val)
}

func ConvInt8Value(val int64) rawtype.Value {
	return rawtype.Value{
		TypeTag:   rawtype.TypeInt8,
		NormValue: bparse.Serialize.Int8(val),
	}
}
