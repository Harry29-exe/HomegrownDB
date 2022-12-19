package toast

const InTupleSize = 18

func IsVarLenToasted(firstByte byte) bool {
	return firstByte == 128
}

func IsCompressed(firstByte byte) bool {
	return firstByte < 128 && firstByte > 63
}

// Status (toast) The Oversized-Attribute Storage Technique status
type Status uint8

const (
	// PLAIN - no out of line storage
	PLAIN Status = iota
	// EXTENDED - allows for compression and out of line storage
	EXTENDED
	// EXTERNAL - allows for out of line storage but not for compression
	EXTERNAL
)
