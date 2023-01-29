package bitmap

type Bitmap struct {
	bytes []byte
}

func New(bits int) Bitmap {
	bytes := bits / 8

	if bytes*8 != bits {
		bytes++
	}
	return Bitmap{bytes: make([]byte, bytes)}
}

func AsBitmap(data []byte) Bitmap {
	return Bitmap{bytes: data}
}

func (b Bitmap) Set(n int) {
	byteNo := n / 8
	bitNo := n - byteNo
	b.bytes[byteNo] = b.bytes[byteNo] | setBitMasks[bitNo]
}

func (b Bitmap) Clear(n int) {
	byteNo := n / 8
	bitNo := n - byteNo
	b.bytes[byteNo] = b.bytes[byteNo] & clearBitMasks[bitNo]
}

func (b Bitmap) ClearAll() {
	for i := 0; i < len(b.bytes); i++ {
		b.bytes[i] = 0
	}
}

func (b Bitmap) IsSet(n int) bool {
	byteNo := n / 8
	bitNo := n - byteNo
	return b.bytes[byteNo]&setBitMasks[bitNo] > 0
}

func (b Bitmap) Bytes() []byte {
	return b.bytes
}

var setBitMasks = [8]byte{
	1, 2, 4, 8,
	16, 32, 64, 128,
}

var clearBitMasks = [8]byte{
	^setBitMasks[0], ^setBitMasks[1], ^setBitMasks[2], ^setBitMasks[3],
	^setBitMasks[4], ^setBitMasks[5], ^setBitMasks[6], ^setBitMasks[7],
}
