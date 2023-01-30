package inputtype

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/hgtype"
)

func ConvInt8(val int64) []byte {
	return bparse.Serialize.Int8(val)
}

func ConvInt8Value(val int64) hgtype.Value {
	return hgtype.Value{
		TypeTag:   hgtype.TypeInt8,
		NormValue: bparse.Serialize.Int8(val),
	}
}
