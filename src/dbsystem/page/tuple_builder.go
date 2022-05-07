package page

import (
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
	"HomegrownDB/io/bparse"
	"bytes"
)

var TupleCreator = tupleCreator{}

type tupleCreator struct{}

// CreateTuple creates new TupleToSave from given columnValues and transaction context,
// Tuple inside is not initialized i.e. it does not have TID (tuple identifier) and ids of
// objects stored outside tuple should be saved to Tuple
func (tc tupleCreator) CreateTuple(tableDef table.Definition, columnValues map[string]any, txContext tx.Context) (TupleToSave, error) {

	sortedValues := make([]any, 0, tableDef.ColumnCount())

	for i := uint16(0); i < tableDef.ColumnCount(); i++ {
		sortedValues[i] = columnValues[tableDef.ColumnName(i)]
	}

	//initialize bitmap
	bitmapLen := tableDef.NullBitmapLen()
	nullBitmap := make([]byte, bitmapLen)
	colCounter := 0
	for currentByte := uint16(0); currentByte < bitmapLen; currentByte++ {
		for bit := uint8(0); bit < 8; bit++ {
			if sortedValues[colCounter] != nil {
				nullBitmap[currentByte] = bparse.Bit.
					SetBit(nullBitmap[currentByte], bit)
			}

			colCounter++
		}
	}

	buffer := bytes.Buffer{}
	var bufferPtr InTuplePtr = 0
	var lobs []LobValueToSave = nil
	var bgValues []BgValueToSave = nil
	for i, colValue := range sortedValues {
		colSerializer := tableDef.ColumnSerializer(column.OrderId(i))
		dataToSave, err := colSerializer.SerializeValue(colValue)
		if err != nil {
			return TupleToSave{}, err
		}

		switch dataToSave.StorePlace() {
		case column.StoreInTuple:
			buffer.Write(dataToSave.Data())
			bufferPtr += uint16(len(dataToSave.Data()))

		case column.StoreInBackground:
			bgValues = append(bgValues, BgValueToSave{
				Value:      dataToSave.Data(),
				PtrInTuple: bufferPtr,
			})
			bufferPtr += dbsystem.BgObjectIdSize

		case column.StoreInLob:
			lobs = append(lobs, LobValueToSave{
				Value:      dataToSave.Data(),
				PtrInTuple: bufferPtr,
			})
			bufferPtr += dbsystem.LobIdSize
		}
	}

	//todo add created_by_tx created/updated by txid amout of command executed
	// and columns
	panic("not implemented")
}

func (tc tupleCreator) name() {

}

type TupleToSave struct {
	Tuple           Tuple
	BgValuesToSave  []BgValueToSave
	LobValuesToSave []LobValueToSave
}

type BgValueToSave struct {
	Value      []byte
	PtrInTuple InTuplePtr // PtrInTuple ptr to place in tuple where should be putted id od bg object
}

type LobValueToSave struct {
	Value      []byte
	PtrInTuple InTuplePtr // PtrInTuple ptr to place in tuple where should be putted id od lob
}
