package typanlr

import (
	"HomegrownDB/backend/new/internal/sqlerr"
	"HomegrownDB/dbsystem/hgtype"
)

func IsAssignable(destType hgtype.TypeData, srcType hgtype.TypeData) error {
	//todo there should be some arguments checking/conversion
	if destType.Tag != srcType.Tag {
		return sqlerr.NewTypeMismatch(destType.Tag, srcType.Tag, nil)
	}

	return nil
}
