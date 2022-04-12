package factory

import (
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/column/types"
	"HomegrownDB/io/bparse"
)

func DeserializeColumnDefinition(serializedData []byte) (col column.Definition, subsequent []byte) {
	deserialized := bparse.NewDeserializer(serializedData)
	columnCode := deserialized.MdString()

	switch columnCode {
	case types.Int2:
		colDef := &types.Int2Column{}
		return colDef, colDef.Deserialize(serializedData)
	default:
		panic("Unknown type to deserialize")
	}
}
