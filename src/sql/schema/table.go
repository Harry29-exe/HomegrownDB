package schema

import (
	"HomegrownDB/io"
	"HomegrownDB/sql/schema/dbtable"
	"errors"
)

type Table interface {
	GetColumn(name string) Column
}

func DeserializeTable(rawData []byte) (Table, error) {
	deserializer := io.NewDeserializer(rawData)
	implName := deserializer.MdString()

	switch implName {
	case dbtable.TableImplName:
		return dbtable.DeserializeTable(rawData), nil
	default:
		return nil, errors.New("no such table implementation")
	}
}

func SerializeTable(table Table) ([]byte, error) {
	switch t := table.(type) {
	case *dbtable.DbTable:
		return dbtable.SerializeTable(*t), nil
	default:
		return nil, errors.New("can not find type of table implementation")
	}
}
