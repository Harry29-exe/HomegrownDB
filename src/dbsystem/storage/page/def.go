package page

import (
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/schema/relation"
)

type Id = uint32

const IdSize = 4

const Size uint16 = dbsystem.PageSize

type Tag struct {
	PageId   Id
	Relation relation.ID
}
