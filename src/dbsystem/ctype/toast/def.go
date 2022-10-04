package toast

const InTupleSize = 18

func IsToasted(firstByte byte) bool {
	return firstByte == 128
}

func IsCompressed(firstByte byte) bool {
	return firstByte < 128 && firstByte > 63
}
