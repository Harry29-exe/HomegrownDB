package sequence

import "HomegrownDB/common/datastructs"

type Sequence[T datastructs.Number] interface {
	Next() T
}
