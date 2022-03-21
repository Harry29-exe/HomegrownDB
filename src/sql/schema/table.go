package schema

import (
	"HomegrownDB/io"
)

type Table struct {
	objectId uint64
	columns  map[string]Column
	name     string
	byteLen  uint32
}

func ReadTable(rawData []byte) *Table {
	tr := tableReader{
		deserializer: io.NewDeserializer(rawData),
		table:        Table{},
	}

	tr.table.objectId = tr.deserializer.Uint64()
	tr.table.name = tr.deserializer.MdString()
	tr.readColumns()

	return &tr.table
}

type tableReader struct {
	deserializer io.Deserializer
	table        Table
}

func (tr *tableReader) readColumns() {
	columnCount := tr.deserializer.Uint16()

	var columnOffset int32 = 0
	var column Column
	for i := uint16(0); i < columnCount; i++ {
		column = tr.readColumn(columnOffset)
		tr.table.columns[column.Name] = column

		if columnOffset > -1 && column.Type.IsFixedSize {
			columnOffset += int32(column.Type.ByteLen)
		} else {
			columnOffset = -1
		}
	}
}

func (tr *tableReader) readColumn(offset int32) Column {
	colName := tr.deserializer.MdString()
	colTypeCode := tr.deserializer.MdString()
	colTypeArgc := tr.deserializer.Uint8()
	var colTypeArgv = make([]int32, colTypeArgc)
	for i := byte(0); i < colTypeArgc; i++ {
		colTypeArgv[i] = tr.deserializer.Int32()
	}

	return Column{
		Name:          colName,
		Type:          GetColumnType(colTypeCode, colTypeArgv),
		Offset:        offset,
		Nullable:      tr.deserializer.Bool(),
		Autoincrement: tr.deserializer.Bool(),
	}
}
