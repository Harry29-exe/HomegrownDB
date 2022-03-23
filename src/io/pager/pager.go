package pager

import (
	"HomegrownDB/sql/schema"
)

type LocalPager[T schema.Table] struct {
	table T
	cache map[uint32]Page[T]
}
