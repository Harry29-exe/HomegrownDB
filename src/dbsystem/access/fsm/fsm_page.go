package fsm

var layerOffsets = []uint16{
	uint16(2 ^ 0 - 1),  // layer 0 offset
	uint16(2 ^ 1 - 1),  // layer 1 offset
	uint16(4 - 1),      // layer 2 offset
	uint16(2 ^ 3 - 1),  // layer 3 offset
	uint16(2 ^ 4 - 1),  // layer 4 offset
	uint16(2 ^ 5 - 1),  // layer 5 offset
	uint16(2 ^ 6 - 1),  // layer 6 offset
	uint16(2 ^ 7 - 1),  // layer 7 offset
	uint16(2 ^ 8 - 1),  // layer 8 offset
	uint16(2 ^ 9 - 1),  // layer 9 offset
	uint16(2 ^ 10 - 1), // layer 10 offset
	uint16(2 ^ 11 - 1), // layer 11 offset
	uint16(2 ^ 12 - 1), // layer 12 offset
	uint16(2 ^ 13 - 1), // layer 13 offset
	uint16(2 ^ 14 - 1), // layer 14 offset
	uint16(2 ^ 15 - 1), // layer 15 offset
}

type page struct {
	layers uint8
	data   []byte
}

func (p page) get(index uint16, layer uint8) byte {
	return p.data[layerOffsets[layer]+index]
}

func (p page) put(index uint16, layer uint8, value uint8) {
	p.data[layerOffsets[layer]+index] = value
}
