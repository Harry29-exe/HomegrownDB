package factory

import (
	"HomegrownDB/sql/schema/column"
	"HomegrownDB/sql/schema/column/types"
)

// todo add errors for definition creation
func CreateDefinition(args column.Args) column.Definition {
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
