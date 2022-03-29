package providers

import (
	"HomegrownDB/io"
	. "HomegrownDB/sql/schema/difinitions"
	"HomegrownDB/sql/schema/impl"
	"errors"
)

func DeserializeTable(rawData []byte) (Table, error) {
	deserializer := io.NewDeserializer(rawData)
	implName := deserializer.MdString()

	switch implName {
	case impl.TableImplName:
		return impl.DeserializeDbTable(rawData), nil
	default:
		return nil, errors.New("no such table implementation")
	}
}

func SerializeTable(table Table) ([]byte, error) {
	switch t := table.(type) {
	case *impl.DbTable:
		return impl.SerializeDbTable(*t), nil
	default:
		return nil, errors.New("can not find type of table implementation")
	}
}
