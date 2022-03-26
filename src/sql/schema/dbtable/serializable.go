package dbtable

import (
	"HomegrownDB/io"
)

func DeserializeDbTable(rawData []byte) *DbTable {
	td := tableDeserializer{
		deserializer: io.NewDeserializer(rawData),
		table:        DbTable{},
	}

	td.table.objectId = td.deserializer.Uint64()
	td.table.name = td.deserializer.MdString()
	td.readColumns()

	return &td.table
}

type tableDeserializer struct {
	deserializer *io.Deserializer
	table        DbTable
}

//todo add null support when calc offset
func (td *tableDeserializer) readColumns() {
	columnCount := td.deserializer.Uint16()

	var columnOffset int32 = 0
	var column *Column
	for i := uint16(0); i < columnCount; i++ {
		column = td.readColumn(columnOffset)
		td.table.columns[column.Name] = column
		td.table.colList = append(td.table.colList, column)

		if columnOffset > -1 && column.Type.LenPrefixSize == 0 {
			columnOffset += int32(column.Type.ByteLen)
		} else {
			columnOffset = -1
		}
	}
}

func (td *tableDeserializer) readColumn(offset int32) *Column {
	colName := td.deserializer.MdString()
	colTypeCode := td.deserializer.MdString()
	colTypeArgc := td.deserializer.Uint8()
	var colTypeArgv = make([]int32, colTypeArgc)
	for i := byte(0); i < colTypeArgc; i++ {
		colTypeArgv[i] = td.deserializer.Int32()
	}

	return &Column{
		Name:          colName,
		Type:          *GetColumnType(colTypeCode, colTypeArgv),
		Offset:        offset,
		Nullable:      td.deserializer.Bool(),
		Autoincrement: td.deserializer.Bool(),
	}
}

type tableSerializer struct {
	serializer *io.Serializer
	table      *DbTable
}

func SerializeDbTable(table DbTable) []byte {
	serializer := io.NewSerializer()
	ts := tableSerializer{
		serializer: serializer,
		table:      &table,
	}

	serializer.Uint64(table.objectId)
	serializer.MdString(table.name)

	columnCount := uint16(len(table.columns))
	serializer.Uint16(columnCount)
	for _, column := range table.columns {
		ts.serializeColumn(column)
	}

	return ts.serializer.GetBytes()
}

func (ts *tableSerializer) serializeColumn(column *Column) {
	ts.serializer.MdString(column.Name)

	columnType := column.Type
	ts.serializer.MdString(columnType.Code)
	ts.serializer.Uint32(columnType.ByteLen)
	ts.serializer.Uint8(columnType.LenPrefixSize)
	ts.serializer.Uint8(columnType.LobStatus)

	ts.serializer.Bool(column.Nullable)
	ts.serializer.Bool(column.Autoincrement)
}
