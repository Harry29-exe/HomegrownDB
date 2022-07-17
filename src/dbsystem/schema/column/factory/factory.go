package factory

import (
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/column/ctypes"
)

// todo add errors for definition creation
func CreateDefinition(args column.Args) column.WDefinition {
	columnType, err := args.Type()
	if err != nil {
		panic("Args must contain column's Type")
	}

	switch columnType {
	case ctypes.Int2:
		return ctypes.NewInt2Column(args)
	default:
		panic("unknown column type: " + columnType)
	}
}
