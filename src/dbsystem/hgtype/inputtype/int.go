package inputtype

import "HomegrownDB/common/bparse"

func ConvInt8(val int64) []byte {
	return bparse.Serialize.Int8(val)
}
