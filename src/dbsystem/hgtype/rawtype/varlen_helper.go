package rawtype

import "HomegrownDB/common/bparse"

var VarLenHelper = varLenHelper{}

type varLenHelper struct{}

func (v varLenHelper) dataLen(value []byte) int {
	if v.lenIsOneByte(value[0]) {
		return v.getOneByteLen(value) - 1
	} else {
		return v.getFourByteLen(value) - 4
	}
}

func (v varLenHelper) fullLen(value []byte) int {
	if v.lenIsOneByte(value[0]) {
		return v.getOneByteLen(value)
	} else {
		return v.getFourByteLen(value)
	}
}

func (v varLenHelper) data(value []byte) []byte {
	return value[v.dataLen(value):]
}

func lenCanBeOneByte(dataLen uint32) bool {
	return dataLen+1 < 128
}

func (v varLenHelper) lenIsOneByte(firstByte byte) bool {
	return firstByte > 127
}

func (v varLenHelper) getFourByteLen(data []byte) int {
	return int(bparse.Parse.UInt4(data) & fourByteHeaderMask)
}

func (v varLenHelper) getOneByteLen(data []byte) int {
	return int(data[0] & oneByteHeaderMask)
}

func toOneByteLen(length uint8) byte {
	return length | (^oneByteHeaderMask)
}
