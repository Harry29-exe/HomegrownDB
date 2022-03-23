package pager

import (
	"HomegrownDB/sql/schema"
)

const PageSize uint32 = 8192

type Page[T schema.Table] struct {
	lastModified uint64
	rows         []byte
}
