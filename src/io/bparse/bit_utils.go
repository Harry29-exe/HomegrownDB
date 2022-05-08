package bparse

var Bit = bitUtils{}

type bitUtils struct{}

// SetBit sets bit at bitInByte position in given byte and returns changed byte
func (b bitUtils) SetBit(byte byte, bitInByte uint8) byte {
	return byte | setBitsMap[bitInByte]
}

// ClearBit clears bit at bitInByte position in given byte and returns changed byte
func (b bitUtils) ClearBit(byte byte, bitInByte uint8) byte {
	return byte &^ setBitsMap[bitInByte]
}

var setBitsMap = [8]byte{
	1, 2, 4, 8,
	16, 32, 64, 128,
}
