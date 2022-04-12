package factory

import (
	column2 "HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/column/types"
)

// todo add errors for definition creation
func CreateDefinition(args column2.Args) column2.Definition {
	columnType, err := args.Type()
	if err != nil {
		panic("Args must contain column's Type")
	}

	switch columnType {
	case types.Int2:
		return types.NewInt2Column(args)
	default:
		panic("unknown column type: " + columnType)
	}
}
