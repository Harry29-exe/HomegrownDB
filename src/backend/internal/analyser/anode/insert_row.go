package anode

type InsertRows struct {
	// data are serialized input values
	data []byte
	// dataPtr stores pointers to starts of values data
	// and pointer which is equal to len(data)
	dataPtr []uint32
	// valuesPerRow number of values that will be stored per row
	valuesPerRow uint32
	// valuesPut number of values that are putted into struct
	valuesPut uint32
}

func NewInsertRows(rowCount uint, valuesPerRow uint16) *InsertRows {
	return &InsertRows{
		dataPtr:      make([]uint32, rowCount*uint(valuesPerRow)+1),
		data:         make([]byte, 0, rowCount*uint(valuesPerRow)*8),
		valuesPerRow: 0,
		valuesPut:    0,
	}
}

func (v *InsertRows) GetValue(rowIndex uint32, valueIndex uint16) []byte {
	valuePtrIndex := rowIndex*v.valuesPerRow + uint32(valueIndex)
	return v.data[v.dataPtr[valuePtrIndex]:v.dataPtr[valuePtrIndex+1]]
}

func (v *InsertRows) PutValue(value []byte) {
	v.data = append(v.data, value...)
	v.valuesPut++
	v.dataPtr[v.valuesPut] = uint32(len(v.data))
}
