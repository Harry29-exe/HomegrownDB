package ctype

import (
	"HomegrownDB/common/bparse"
	"strconv"
)

func convInputInt(input int64, cType CType) ([]byte, error) {
	switch cType {
	case Int8:
		return bparse.Serialize.Int8(input), nil
	case Float8:
		float := float64(input)
		return bparse.Serialize.Float8(float), nil
	default:
		return nil, NewInputConvErr(InputTypeInt, strconv.Itoa(int(input)), cType)
	}
}
