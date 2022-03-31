package table

import (
	"HomegrownDB/io/bparse"
)

func DeserializeDbTable(rawData []byte) *DBTable {
	td := tableDeserializer{
		deserializer: bparse.NewDeserializer(rawData),
		table:        DBTable{},
	}

	td.table.objectId = td.deserializer.Uint64()
	td.table.name = td.deserializer.MdString()
	td.readColumns()

	return &td.table
}

type tableDeserializer struct {
	deserializer *bparse.Deserializer
	table        DBTable
}

//todo add null support when calc offset
func (td *tableDeserializer) readColumns() {
	//columnCount := td.deserializer.Uint16()
	//
	//var columnOffset int32 = 0
	//var column *definitions.Column
	//for i := uint16(0); i < columnCount; i++ {
	//	column = td.readColumn(columnOffset)
	//	td.table.colNameIdMap[column.Name] = column
	//	td.table.colList = append(td.table.colList, column)
	//
	//	if columnOffset > -1 && column.Type.LenPrefixSize == 0 {
	//		columnOffset += int32(column.Type.ByteLen)
	//	} else {
	//		columnOffset = -1
	//	}
	//}
	panic("")
}
