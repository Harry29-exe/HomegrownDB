package pager

import "HomegrownDB/sql"

type LocalPager[T sql.Table] struct {
	table T
	cache map[uint32]Page[T]
}
