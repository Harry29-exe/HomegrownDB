package typanlr

import (
	"HomegrownDB/backend/internal/sqlerr"
	"HomegrownDB/dbsystem/hgtype"
)

func IsAssignable(destType hgtype.ColumnType, srcType hgtype.ColumnType) error {
	//todo there should be some arguments checking/conversion
	if destType.ColTag != srcType.ColTag {
		return sqlerr.NewTypeMismatch(destType.ColTag, srcType.ColTag, nil)
	}

	return nil
}
