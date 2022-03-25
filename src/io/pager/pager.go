package pager

import (
	"HomegrownDB/sql/schema/dbtable"
)

type LocalPager[T dbtable.DbTable] struct {
	table T
	cache map[uint32]Page[T]
}
