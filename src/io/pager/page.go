package pager

import "HomegrownDB/sql"

const PageSize uint32 = 8192

type Page[T sql.Table] struct {
	lastModified uint64
	rows         []byte
}
