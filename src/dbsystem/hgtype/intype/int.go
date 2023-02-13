package intype

import (
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/lib/bparse"
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
