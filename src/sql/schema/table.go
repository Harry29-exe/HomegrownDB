package schema

import (
	"HomegrownDB/io"
	"HomegrownDB/sql/schema/dbtable"
	"errors"
)

type Table interface {
	GetColumn(name string) Column
	ParseRow(row []byte) ParsedRow

	ParseColumn(columnId uint16, row []byte) []byte
	ParseColumns(columnsIDs []uint16, row []byte) [][]byte
}

type ParsedRow interface {
	GetColumn(id int16) []byte
}

func DeserializeTable(rawData []byte) (Table, error) {
	deserializer := io.NewDeserializer(rawData)
	implName := deserializer.MdString()

	switch implName {
	case dbtable.TableImplName:
		return dbtable.DeserializeDbTable(rawData), nil
	default:
		return nil, errors.New("no such table implementation")
	}
}

func SerializeTable(table Table) ([]byte, error) {
	switch t := table.(type) {
	case *dbtable.DbTable:
		return dbtable.SerializeDbTable(*t), nil
	default:
		return nil, errors.New("can not find type of table implementation")
	}
}
