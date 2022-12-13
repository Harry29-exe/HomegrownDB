package ctype

import (
	"HomegrownDB/common/bparse"
	"strconv"
)

func convInputInt(input int64, cType Type) ([]byte, error) {
	switch cType {
	case TypeInt8:
		return bparse.Serialize.Int8(input), nil
	case TypeFloat8:
		float := float64(input)
		return bparse.Serialize.Float8(float), nil
	default:
		return nil, NewInputConvErr(InputTypeInt, strconv.Itoa(int(input)), cType)
	}
}
