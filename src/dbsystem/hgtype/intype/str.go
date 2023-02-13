package intype

import (
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/lib/bparse"
	"fmt"
)

func ConvStr(val string) ([]byte, error) {
	l := len(val) + 4
	if l > rawtype.MaxVarLenSize {
		return nil, fmt.Errorf("string values can not be longer that %d", rawtype.MaxVarLenSize) //todo better err
	}
	serializedVal := make([]byte, l)
	bparse.Serialize.PutUInt4(uint32(l), serializedVal)
	copy(serializedVal[4:], val)
	return serializedVal, nil
}

func ConvStrValue(val string) (rawtype.Value, error) {
	normValue, err := ConvStr(val)
	if err != nil {
		return rawtype.Value{}, err
	}
	return rawtype.Value{
		TypeTag:   rawtype.TypeStr,
		NormValue: normValue,
	}, err
}
