package typanlr

import (
	"HomegrownDB/backend/internal/sqlerr"
	"HomegrownDB/dbsystem/hgtype"
)

func IsAssignable(destType hgtype.ColumnType, srcType hgtype.ColumnType) error {
	//todo there should be some arguments checking/conversion
	if destType.Tag != srcType.Tag {
		return sqlerr.NewTypeMismatch(destType.Tag, srcType.Tag, nil)
	}

	return nil
}
