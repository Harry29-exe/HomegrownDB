package factory

import (
	"HomegrownDB/io/bparse"
	"HomegrownDB/sql/schema/column"
	"HomegrownDB/sql/schema/column/types"
)

func DeserializeColumnDefinition(serializedData []byte) column.Definition {
	deserialized := bparse.NewDeserializer(serializedData)
	columnCode := deserialized.MdString()

	switch columnCode {
	case types.Int2:
		colDef := &types.Int2Column{}
		colDef.Deserialize(serializedData)
		return colDef
	default:
		panic("Unknown type to deserialize")
	}
}
