package factory

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/column/ctypes"
)

func DeserializeColumnDefinition(serializedData []byte) (col column.WDefinition, subsequent []byte) {
	deserialized := bparse.NewDeserializer(serializedData)
	columnCode := deserialized.MdString()

	switch columnCode {
	case ctypes.Int2:
		colDef := &ctypes.Int2Column{}
		return colDef, colDef.Deserialize(serializedData)
	default:
		panic("Unknown type to deserialize")
	}
}
