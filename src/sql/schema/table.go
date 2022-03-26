package schema

import (
	"HomegrownDB/io"
	"HomegrownDB/sql/schema/dbtable"
	"errors"
)

type Table interface {
	GetColumnId(name string) ColumnId
	GetColumnsIds(names []string) []ColumnId
	ParseRow(row []byte) Tuple

	RetrieveColumn(columnId ColumnId, row []byte) TupleColumn
	RetrieveColumns(columnsIDs []ColumnId, row []byte) []TupleColumn
}

type ColumnId = uint16

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
