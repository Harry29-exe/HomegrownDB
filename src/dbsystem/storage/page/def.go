package page

import (
	"HomegrownDB/dbsystem"
	"math"
)

type Id = uint32

const IdSize = 4

const (
	MaxId     = math.MaxUint32 - 1
	InvalidId = math.MaxUint32
)

const Size uint16 = dbsystem.PageSize
