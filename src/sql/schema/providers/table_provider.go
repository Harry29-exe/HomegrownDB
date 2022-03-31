package providers

import (
	"HomegrownDB/io/bparse"
	"HomegrownDB/sql/schema/table"
	"errors"
)

func DeserializeTable(rawData []byte) (table.Table, error) {
	deserializer := bparse.NewDeserializer(rawData)
	implName := deserializer.MdString()

	switch implName {
	case table.TableImplName:
		return table.DeserializeDbTable(rawData), nil
	default:
		return nil, errors.New("no such table implementation")
	}
}

func SerializeTable(table table.Table) ([]byte, error) {
	switch t := table.(type) {
	case *table.DbTable:
		return table.SerializeDbTable(*t), nil
	default:
		return nil, errors.New("can not find type of table implementation")
	}
}
