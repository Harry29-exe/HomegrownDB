package anode

type Values struct {
	// data are serialized input values
	data []byte
	// dataPtr stores pointers to starts of values data
	// and pointer which is equal to len(data)
	dataPtr []uint32
	// valuesSize number of values that will be stored
	valuesSize uint16
	// valuesPut number of values that are putted into struct
	valuesPut uint16
}

func NewValues(valuesSize uint16) *Values {
	return &Values{
		data:       make([]byte, valuesSize*6),
		dataPtr:    make([]uint32, valuesSize+1),
		valuesSize: valuesSize,
	}
}

func (v *Values) GetValue(index int) []byte {
	return v.data[v.dataPtr[index]:v.dataPtr[index+1]]
}

func (v *Values) PutValue(value []byte) {
	if v.valuesPut == v.valuesSize {
		panic("illegal state")
	}
	bytesCopied := copy(v.data[v.dataPtr[v.valuesPut]:], value)
	v.dataPtr[v.valuesPut+1] = v.dataPtr[v.valuesPut] + uint32(bytesCopied)
	v.valuesPut++
}
